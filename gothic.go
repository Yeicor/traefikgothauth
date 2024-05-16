package traefikgothauth

import (
	"errors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/http"
	"net/url"
)

// HACK: gothic logouts automatically during each login, so we copy the code without the logout.

var CompleteUserAuthNoLogout = func(res http.ResponseWriter, req *http.Request) (goth.User, error) {
	providerName, err := gothic.GetProviderName(req)
	if err != nil {
		return goth.User{}, err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}

	value, err := gothic.GetFromSession(providerName, req)
	if err != nil {
		return goth.User{}, err
	}
	//HACK: Removed defer gothic.Logout(res, req)
	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// HACK: validateState only after FetchUser fails
	err = validateState(req, sess)
	if err != nil {
		return goth.User{}, err
	}

	params := req.URL.Query()
	if params.Encode() == "" && req.Method == "POST" {
		_ = req.ParseForm()
		params = req.Form
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, params)
	if err != nil {
		return goth.User{}, err
	}

	err = gothic.StoreInSession(providerName, sess.Marshal(), req, res)

	if err != nil {
		return goth.User{}, err
	}

	gu, err := provider.FetchUser(sess)
	return gu, err
}

func validateState(req *http.Request, sess goth.Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	reqState := gothic.GetState(req)

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != reqState) {
		return errors.New("state token mismatch")
	}
	return nil
}
