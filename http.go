package traefikgothauth

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
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
	req = gothic.GetContextWithProvider(req, o.config.ProviderName)
	req = mux.SetURLVars(req, map[string]string{}) // Avoid yaegi bug

	// Handle logout requests.
	if o.config.LogoutURI != "" && req.URL.Path == o.config.logoutURI.Path {
		logd("Logging out")
		err := gothic.Logout(rw, req)
		if err != nil {
			loge("Failed to logout", "error", err)
			http.Error(rw, "Failed to logout", http.StatusInternalServerError)
			return
		}
		http.ServeContent(rw, req, "index.html", time.Now(), strings.NewReader("Logged out successfully!"))
		return
	}

	// Detect unauthenticated requests and begin the authentication process.
	logd("Checking authentication")
	auth, err := CompleteUserAuthNoLogout(rw, req)
	if err != nil {
		if req.URL.Path == o.config.providerCallback.Path {
			loge("Failed to authenticate", "error", err)
			http.Error(rw, "Failed to authenticate", http.StatusInternalServerError)
		} else {
			// NOTE: Handling them here avoids possible infinite loop when failing after redirecting to the
			// providerCallback url as that matches the previous if statement.
			o.runBeginAuthHandler(rw, req)
		}
		return
	}

	// Handle the callback from the provider, redirect to the initial URL after login success.
	if req.URL.Path == o.config.providerCallback.Path {
		// We authenticated successfully, check authorization!
		if o.config.Authorize != nil {
			fillRawData(&auth)
			for key, regex := range o.config.Authorize.regexes {
				if !regex.MatchString(fmt.Sprint(auth.RawData[key])) {
					loge("Unauthorized user tried to log in", "bad-key", key, "bad-value", fmt.Sprint(auth.RawData[key]))
					// If authorization fails, never send the cookie to the client as it gives unconditional access!
					rw.Header().Del("Set-Cookie")
					http.Error(rw, "Unauthorized user", http.StatusForbidden)
					return
				}
			}
		}
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
		logi("User just logged in!", "user", fmt.Sprintf("%+v", auth.RawData), "redirect", redirectPath)
		http.Redirect(rw, req, redirectPath, http.StatusTemporaryRedirect)
		return
	}

	// We are authenticated with this provider, publish claims (if enabled) and finish!
	if o.config.ClaimsPrefix != "__NO__" {
		fillRawData(&auth)
		logt("Publishing claims for next http handler", "claims", fmt.Sprintf("%+v", auth.RawData))
		for key, value := range auth.RawData {
			headerKey := http.CanonicalHeaderKey(o.config.ClaimsPrefix + invalidHeader.ReplaceAllString(key, "-"))
			req.Header.Add(headerKey, fmt.Sprintf("%v", value))
		}
	}

	// Authentication completed, run the next handler.
	o.next.ServeHTTP(rw, req)
}

func (o *Plugin) runBeginAuthHandler(rw http.ResponseWriter, req *http.Request) {
	logd("Starting new authentication")
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
	//auth.RawData["access-token-secret"] = auth.AccessTokenSecret
	//auth.RawData["refresh-token"] = auth.RefreshToken
	auth.RawData["expires-at"] = auth.ExpiresAt
	//auth.RawData["id-token"] = auth.IDToken
	// Drop empty values
	for key, value := range auth.RawData {
		if value == "" {
			delete(auth.RawData, key)
		}
	}
}

var invalidHeader = regexp.MustCompile("[^a-zA-Z0-9-]+") // Also removing _ from headers
