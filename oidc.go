// Package traefikoidc is a Traefik plugin to authenticate requests using OpenID Connect.
package traefikoidc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Config configures the OpenID Connect plugin.
type Config struct {
	// AuthorizationEndpoint is the OpenID Connect server's authorization endpoint (get it from .../.well-known/openid-configuration).
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// TokenEndpoint is the OpenID Connect server's token endpoint (get it from .../.well-known/openid-configuration).
	TokenEndpoint string `json:"token_endpoint"`
	// ClientID is the OAuth2 client ID.
	ClientID string `json:"client_id"`
	// ClientSecret is the OAuth2 client secret.
	ClientSecret string `json:"client_secret"`
	// RedirectURL is the OAuth2 client redirect URL. Can be a relative path!
	RedirectURL string `json:"redirect_url"`
	// Scopes is a space-separated string with the OAuth2 scopes.
	Scopes string `json:"scopes"`
	// Cookie is the name of the authentication cookie.
	Cookie string `json:"cookie"`
	// ClaimsPrefix is the prefix for the claims to be published as headers.
	ClaimsPrefix string `json:"claims_prefix"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		AuthorizationEndpoint: "https://accounts.google.com/o/oauth2/v2/auth",
		RedirectURL:           "/oidc/callback/",
		Scopes:                "openid",
		Cookie:                "oidc-auth",
		ClaimsPrefix:          "x-auth-",
	}
}

// OIDC is the OpenID Connect plugin.
type OIDC struct {
	next        http.Handler
	name        string
	config      *Config
	redirectURL *url.URL
}

// New created a new OIDC plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	redirectURL, err := url.Parse(config.RedirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redirect URL %q: %v", config.RedirectURL, err)
	}

	return &OIDC{
		next:        next,
		name:        name,
		config:      config,
		redirectURL: redirectURL,
	}, nil
}

func (o *OIDC) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Handle callback requests.
	if req.URL.Path == o.redirectURL.Path {
		if err := o.handleOAuth2Callback(rw, req); err != nil {
			loge("Failed to handle OAuth2 callback", "error", err)
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	// If the authentication cookie doesn't exist, redirect to OAuth2 provider.
	authCookie, err := req.Cookie(o.config.Cookie)
	if err != nil {
		logd("Cookie not found", "error", err, "cookie_name", o.config.Cookie, "cookies", req.Cookies())
		o.handleOAuth2Redirect(rw, req)
		return
	}

	// Check if the cookie is valid, or redirect to OAuth2 provider to refresh it.
	idToken, err := oauth2Verify(authCookie.Value)
	if err != nil {
		logd("Cookie is invalid (expired?)", "error", err)
		o.handleOAuth2Redirect(rw, req)
		return
	}

	// Extract all claims and publish them as request headers for the next handlers.
	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		loge("Failed to extract claims", "error", err)
		o.handleOAuth2Redirect(rw, req)
		return
	}
	for key, value := range claims {
		headerKey := http.CanonicalHeaderKey(o.config.ClaimsPrefix + invalidHeader.ReplaceAllString(key, "-"))
		req.Header.Add(headerKey, fmt.Sprintf("%v", value))
	}

	// Authentication completed, continue processing.
	o.next.ServeHTTP(rw, req)
}

func (o *OIDC) handleOAuth2Redirect(rw http.ResponseWriter, req *http.Request) {
	// Redirect to OAuth2 provider.
	authUrl, err := url.Parse(o.config.AuthorizationEndpoint)
	if err != nil {
		loge("Failed to parse authorization endpoint", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	state, err := o.encodeState(req.URL)
	if err != nil {
		loge("Failed to encode state", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	redirectUri := o.redirectUriFor(req).String()
	query := authUrl.Query()
	query.Set("response_type", "code")
	query.Set("client_id", o.config.ClientID)
	query.Set("redirect_uri", redirectUri)
	query.Set("scope", o.config.Scopes)
	query.Set("state", state)
	authUrl.RawQuery = query.Encode()
	logd("Redirecting to OAuth2 provider", "auth_url", authUrl, "redirect_uri", redirectUri, "state", state, "state_url", req.URL)
	http.Redirect(rw, req, authUrl.String(), http.StatusFound)
}

var invalidHeader = regexp.MustCompile("[^a-zA-Z0-9-]+") // Also removing _ from headers
func (o *OIDC) handleOAuth2Callback(w http.ResponseWriter, req *http.Request) error {
	logd("Handling OAuth2 callback", "query", req.URL.Query())

	// Verify state and errors.
	rawIDToken, err := o.oAuth2Exchange(req, req.URL.Query().Get("code"))
	if err != nil {
		return fmt.Errorf("failed to exchange token: %w", err)
	}

	// Parse and verify ID Token payload.
	idToken, err := oauth2Verify(rawIDToken)
	if err != nil {
		return fmt.Errorf("failed to verify ID Token: %w", err)
	}

	logi("User authenticated", "subject", idToken.Subject, "expiry", idToken.Expiry)

	// Set the authentication cookie, and continue processing.
	cookie := &http.Cookie{
		Name:     o.config.Cookie,
		Value:    rawIDToken,
		Expires:  idToken.expiry(),
		Path:     "/",                  // Allow all paths to read the cookie.
		SameSite: http.SameSiteLaxMode, // Allow all subdomains to read the cookie.
		HttpOnly: true,                 // Prevent JavaScript from reading the cookie.
	}
	req.AddCookie(cookie)     // Fake the cookie for the current request to avoid another round-trip.
	http.SetCookie(w, cookie) // Set the cookie in the response for future requests.

	// Redirect to the original request.
	newURL, err := o.decodeState(req.URL.Query().Get("state"))
	if err != nil {
		return fmt.Errorf("failed to decode state: %w", err)
	}

	http.Redirect(w, req, newURL.String(), http.StatusFound)
	return nil
}

func (o *OIDC) oAuth2Exchange(req *http.Request, code string) (string, error) {
	// Exchange the code for a token.
	tokenUrl, err := url.Parse(o.config.TokenEndpoint)
	if err != nil {
		return "", fmt.Errorf("failed to parse token endpoint: %w", err)
	}
	query := tokenUrl.Query()
	query.Set("grant_type", "authorization_code")
	query.Set("code", code)
	query.Set("redirect_uri", o.redirectUriFor(req).String())
	query.Set("client_id", o.config.ClientID)
	query.Set("client_secret", o.config.ClientSecret)
	tokenUrl.RawQuery = query.Encode()
	logd("Exchanging code for token", "token_url", tokenUrl)
	resp, err := http.Post(tokenUrl.String(), "application/x-www-form-urlencoded", nil)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			loge("Failed to close response body", "error", err)
		}
	}(resp.Body)
	var tokenResponse struct {
		IDToken string `json:"id_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}
	if tokenResponse.IDToken == "" {
		return "", fmt.Errorf("token response is empty")
	}
	return tokenResponse.IDToken, nil
}

func oauth2Verify(token string) (idToken, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return idToken{}, fmt.Errorf("invalid token format")
	}
	// FIXME: Verify the signature of the token.
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return idToken{}, fmt.Errorf("failed to decode token payload: %w", err)
	}
	var idToken idToken
	if err := json.Unmarshal(payload, &idToken); err != nil {
		return idToken, fmt.Errorf("failed to unmarshal token payload: %w", err)
	}
	return idToken, nil
}

func (o *OIDC) redirectUriFor(req *http.Request) *url.URL {
	myRedirectURL := o.redirectURL
	if myRedirectURL.Host == "" { // In case no host is specified, fill it with the current request's host.
		if req.TLS != nil {
			myRedirectURL.Scheme = "https"
		} else {
			myRedirectURL.Scheme = "http"
		}
		myRedirectURL.Host = req.Host
	}
	return myRedirectURL
}

func (o *OIDC) encodeState(url *url.URL) (string, error) {
	return base64.URLEncoding.EncodeToString([]byte(url.String())), nil
}

func (o *OIDC) decodeState(state string) (*url.URL, error) {
	bs, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return nil, err
	}
	return url.Parse(string(bs))
}

type idToken struct {
	Issuer       string                 `json:"iss"`
	Subject      string                 `json:"sub"`
	Audience     string                 `json:"aud"`
	Expiry       int64                  `json:"exp"`
	IssuedAt     int64                  `json:"iat"`
	NotBefore    *int64                 `json:"nbf"`
	Nonce        string                 `json:"nonce"`
	AtHash       string                 `json:"at_hash"`
	ClaimNames   map[string]string      `json:"_claim_names"`
	ClaimSources map[string]claimSource `json:"_claim_sources"`
}

func (t *idToken) expiry() time.Time {
	return time.Unix(t.Expiry, 0)
}

func (t *idToken) Claims(m *map[string]interface{}) error {
	return nil // TODO: extract all claims from the token.
}

type claimSource struct {
	Endpoint    string `json:"endpoint"`
	AccessToken string `json:"access_token"`
}
