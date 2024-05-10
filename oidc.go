// Package traefikoidc is a Traefik plugin to authenticate requests using OpenID Connect.
package traefikoidc

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Config configures the OpenID Connect plugin.
type Config struct {
	// ProviderURL is the OpenID Connect server's base URL (needs to have a .well-known/openid-configuration endpoint).
	ProviderURL string `json:"provider_url"`
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
		ProviderURL:  "https://accounts.google.com",
		RedirectURL:  "/oidc/callback/",
		Scopes:       oidc.ScopeOpenID,
		Cookie:       "oidc-auth",
		ClaimsPrefix: "x-auth-",
	}
}

// OIDC is the OpenID Connect plugin.
type OIDC struct {
	next        http.Handler
	name        string
	config      *Config
	provider    *oidc.Provider
	redirectURL *url.URL
	verifier    *oidc.IDTokenVerifier
}

// New created a new OIDC plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	provider, err := oidc.NewProvider(ctx, config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to query provider %q: %v", config.ProviderURL, err)
	}

	redirectURL, err := url.Parse(config.RedirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redirect URL %q: %v", config.RedirectURL, err)
	}

	return &OIDC{
		next:        next,
		name:        name,
		config:      config,
		provider:    provider,
		redirectURL: redirectURL,
		verifier:    provider.Verifier(&oidc.Config{ClientID: config.ClientID}),
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
	idToken, err := o.verifier.Verify(req.Context(), authCookie.Value)
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
	config := o.oAuth2Config(req)
	state, err := o.encodeState(req.URL)
	if err != nil {
		loge("Failed to encode state", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	logd("Redirecting to OAuth2 provider", "redirect_url", config.RedirectURL, "state", state, "state_url", req.URL)
	http.Redirect(rw, req, config.AuthCodeURL(state), http.StatusFound)
}

var invalidHeader = regexp.MustCompile("[^a-zA-Z0-9-]+") // Also removing _ from headers
func (o *OIDC) handleOAuth2Callback(w http.ResponseWriter, req *http.Request) error {
	logd("Handling OAuth2 callback", "query", req.URL.Query())

	// Verify state and errors.
	oauth2Token, err := o.oAuth2Config(req).Exchange(req.Context(), req.URL.Query().Get("code"))
	if err != nil {
		return fmt.Errorf("failed to exchange token: %w", err)
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return fmt.Errorf("no id_token in token response")
	}

	// Parse and verify ID Token payload.
	idToken, err := o.verifier.Verify(req.Context(), rawIDToken)
	if err != nil {
		return fmt.Errorf("failed to verify ID Token: %w", err)
	}

	logi("User authenticated", "subject", idToken.Subject, "expiry", idToken.Expiry)

	// Set the authentication cookie, and continue processing.
	cookie := &http.Cookie{
		Name:     o.config.Cookie,
		Value:    rawIDToken,
		Expires:  idToken.Expiry,
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

func (o *OIDC) oAuth2Config(req *http.Request) *oauth2.Config {
	myRedirectURL := o.redirectURL
	if myRedirectURL.Host == "" { // In case no host is specified, fill it with the current request's host.
		if req.TLS != nil {
			myRedirectURL.Scheme = "https"
		} else {
			myRedirectURL.Scheme = "http"
		}
		myRedirectURL.Host = req.Host
	}

	return &oauth2.Config{
		Endpoint:     o.provider.Endpoint(),
		ClientID:     o.config.ClientID,
		ClientSecret: o.config.ClientSecret,
		RedirectURL:  myRedirectURL.String(),
		Scopes:       strings.Split(o.config.Scopes, " "),
	}
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
