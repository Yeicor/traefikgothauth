package traefikgothauth

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func (o *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if logtEnabled() {
		session, _ := gothic.Store.Get(req, gothic.SessionName)
		sessionKeys := make([]string, 0, len(session.Values))
		for key := range session.Values {
			sessionKeys = append(sessionKeys, fmt.Sprint(key))
		}
		logt("Request", "method", req.Method, "url", req.URL.String(), "remote", req.RemoteAddr, "sessionKeys", sessionKeys)
	}
	for _, providerConfig := range o.config.Providers {
		req = gothic.GetContextWithProvider(req, providerConfig.Name)

		// Handle logout requests.
		if req.URL.Path == providerConfig.logoutURI.Path {
			logd("Logging out", "provider", providerConfig.Name)
			err := gothic.Logout(rw, req)
			if err != nil {
				loge("Failed to logout", "provider", providerConfig.Name, "error", err)
				http.Error(rw, "Failed to logout", http.StatusInternalServerError)
				return
			}
			http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
			return
		}

		// Handle callback/redirect_uri requests, and normal requests that are already authenticated.
		logd("Completing authentication", "provider", providerConfig.Name)
		auth, err := CompleteUserAuthNoLogout(rw, req)
		if err != nil {
			if req.URL.Path == providerConfig.redirectURI.Path {
				loge("Failed to authenticate", "provider", providerConfig.Name, "error", err)
				http.Error(rw, "Failed to authenticate", http.StatusInternalServerError)
				return
			} else {
				logd("Not authenticated", "provider", providerConfig.Name, "error", err)
				// Handle login requests that specify the providerConfig.
				// NOTE: Handling them here avoids possible infinite loop when redirecting to the login url
				if req.URL.Path == providerConfig.authURI.Path {
					o.runBeginAuthHandler(rw, req, providerConfig)
					return
				}
				continue
			}
		}
		if req.URL.Path == providerConfig.redirectURI.Path {
			auth.IDToken = strings.Repeat("*", len(auth.IDToken))
			auth.AccessToken = strings.Repeat("*", len(auth.AccessToken))
			auth.AccessTokenSecret = strings.Repeat("*", len(auth.AccessTokenSecret))
			auth.RefreshToken = strings.Repeat("*", len(auth.RefreshToken))
			// Redirect to initial URL after login success!
			redirectPath := "/" // Default if it cannot be recovered
			redirectSession, err := gothic.Store.Get(req, gothic.SessionName+"_redirect")
			if err == nil {
				redirectPathTmp, ok := redirectSession.Values["path"].(string)
				if ok {
					redirectPath = redirectPathTmp
				} else {
					err = errors.New("could not get the path value from the redirect session cookie")
				}
			}
			if err != nil {
				logw("Could not recover the redirect path", "error", err.Error())
			}
			logi("User just logged in", "provider", providerConfig.Name, "user", fmt.Sprintf("%+v", auth), "redirect", redirectPath)
			http.Redirect(rw, req, redirectPath, http.StatusTemporaryRedirect)
			return
		}

		// We are authenticated with this provider, publish claims and finish!
		fillRawData(&auth)
		logt("Publishing claims for next http handler", "provider", providerConfig.Name, "claims", fmt.Sprintf("%+v", auth.RawData))
		for key, value := range auth.RawData {
			headerKey := http.CanonicalHeaderKey(o.config.ClaimsPrefix + invalidHeader.ReplaceAllString(key, "-"))
			req.Header.Add(headerKey, fmt.Sprintf("%v", value))
		}

		// Authentication completed, run the next handler.
		o.next.ServeHTTP(rw, req)
		return
	}

	// We could not authenticate with any provider, select one to start the authentication.
	var autoBeginAuthFor *ProviderConfig
	if len(o.config.Providers) == 1 {
		autoBeginAuthFor = o.config.Providers[0]
	} else if req.URL.Query().Has("provider") {
		search := req.URL.Query().Get("provider")
		for _, provider := range o.config.Providers {
			if search == provider.Name {
				autoBeginAuthFor = provider
				break
			}
		}
		if autoBeginAuthFor == nil {
			loge("Provider not found", "provider", search)
			http.Error(rw, "Invalid provider", http.StatusBadRequest)
			return
		}
	}
	//autoBeginAuthFor = nil
	//o.providersInfo = allProviders
	if autoBeginAuthFor != nil {
		// Log in with the selected provider without an intermediate page
		o.runBeginAuthHandler(rw, req, autoBeginAuthFor)
	} else {
		// Show a page for the user to choose the provider
		if loginChooseProviderHtmlCache == nil {
			tmp := &bytes.Buffer{}
			err := loginChooseProviderHtml.Execute(tmp, o.providersInfo)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			loginChooseProviderHtmlCache = bytes.NewReader(tmp.Bytes())
			loginChooseProviderHtmlCacheTime = time.Now()
		}
		http.ServeContent(rw, req, "login-choose-provider.html", loginChooseProviderHtmlCacheTime, loginChooseProviderHtmlCache)
	}
}

func (o *Plugin) runBeginAuthHandler(rw http.ResponseWriter, req *http.Request, providerConfig *ProviderConfig) {
	logd("Authenticating", "provider", providerConfig.Name)
	req = gothic.GetContextWithProvider(req, providerConfig.Name)
	redirectSession, err := o.redirectStore.New(req, gothic.SessionName+"_redirect")
	if err == nil {
		redirectSession.Values["path"] = req.RequestURI
		err = redirectSession.Save(req, rw)
	}
	if err != nil {
		logw("Could not save redirect path", "error", err.Error())
	}
	gothic.BeginAuthHandler(rw, req)
}

func fillRawData(auth *goth.User) {
	if auth.RawData == nil {
		auth.RawData = make(map[string]interface{})
	}
	auth.RawData["provider"] = auth.Provider
	auth.RawData["email"] = auth.Email
	auth.RawData["name"] = auth.Name
	auth.RawData["first-name"] = auth.FirstName
	auth.RawData["last-name"] = auth.LastName
	auth.RawData["nick-name"] = auth.NickName
	auth.RawData["description"] = auth.Description
	auth.RawData["user-id"] = auth.UserID
	auth.RawData["avatar-url"] = auth.AvatarURL
	auth.RawData["location"] = auth.Location
	//auth.RawData["access-token"] = auth.AccessToken
	//auth.RawData["refresh-token"] = auth.RefreshToken
	auth.RawData["expires-at"] = auth.ExpiresAt
	// Drop empty values
	for key, value := range auth.RawData {
		if value == "" {
			delete(auth.RawData, key)
		}
	}
}

var invalidHeader = regexp.MustCompile("[^a-zA-Z0-9-]+") // Also removing _ from headers

var loginChooseProviderHtmlCache *bytes.Reader
var loginChooseProviderHtmlCacheTime time.Time
