package traefikgothauth

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/url"
	"strings"
)

// Config configures the Goth Auth plugin.
type Config struct {
	// Providers is the list of configured providers.
	Providers []*ProviderConfig
	// CookieSecret is the secret used to sign the cookie.
	CookieSecret string
	// CookieOptions are the cookie options.
	CookieOptions *sessions.Options
	// ClaimsPrefix is the prefix for the claims to be published as headers.
	ClaimsPrefix string
	// LogLevel is the log level (trace, debug, info, warn, error, off).
	LogLevel string
}

type ProviderConfig struct {
	// Name is the internal name of the provider. There should be only one instance per middleware with a given name.
	Name string
	// ClientKey is the client key for the provider.
	ClientKey string
	// Secret is the secret for the provider.
	Secret string
	// RedirectUri is the full redirect URI for the provider, including the host.
	RedirectURI string
	redirectURI *url.URL
	// AuthURI (optional) is the URI to authenticate against the provider.
	AuthURI string
	authURI *url.URL
	// LogoutURI (optional) is the URI to logout from the provider.
	LogoutURI string
	logoutURI *url.URL
	// Scopes (optional) is the list of scopes for the provider.
	Scopes []string
	// Custom (optional) is the custom configuration for the provider.
	Custom map[string]interface{}
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		CookieOptions: &sessions.Options{HttpOnly: true, Path: "/", MaxAge: 60 * 60},
		ClaimsPrefix:  "X-Auth-",
		LogLevel:      "info",
	}
}

func (c *Config) setup() ([]*ProviderInfo, error) {
	var ok bool
	logLevelCurrent, ok = logTextLevel[strings.ToUpper(c.LogLevel)]
	if !ok {
		loge("Invalid log level", "level", c.LogLevel)
		logLevelCurrent = logLevelInfo
	}
	// TODO: Can this global store cause conflicts between multiple plugin instances?
	//  If this is a problem a possible fix is to rewrite the (small) gothic package
	//  The same happens for goth.UseProviders unless we append a prefix to each one...
	//  The same happens for logLevelCurrent...
	gothic.Store = sessions.NewCookieStore([]byte(c.CookieSecret))
	gothic.Store.(*sessions.CookieStore).Options = c.CookieOptions
	providersInfo := make([]*ProviderInfo, 0, len(c.Providers))
	for _, providerConfig := range c.Providers {
		if providerConfig.RedirectURI == "" {
			return nil, fmt.Errorf("I will not guess your domain name, so you must specify the redirect URI as configured for your provider %s", providerConfig.Name)
		}
		var err error
		providerConfig.redirectURI, err = url.Parse(providerConfig.RedirectURI)
		if err != nil {
			return nil, fmt.Errorf("failed to parse redirect URI: %w", err)
		}
		if providerConfig.redirectURI.Host == "" {
			return nil, fmt.Errorf("redirect URI must include the host: %s", providerConfig.RedirectURI)
		}
		if providerConfig.AuthURI == "" {
			providerConfig.AuthURI = "/__goth/" + providerConfig.Name + "/login/"
		}
		providerConfig.authURI, err = url.Parse(providerConfig.AuthURI)
		if err != nil {
			return nil, fmt.Errorf("failed to parse default auth URI: %w", err)
		}
		if providerConfig.LogoutURI == "" {
			providerConfig.LogoutURI = "/__goth/" + providerConfig.Name + "/logout/"
		}
		providerConfig.logoutURI, err = url.Parse(providerConfig.LogoutURI)
		if err != nil {
			return nil, fmt.Errorf("failed to parse default logout URI: %w", err)
		}
		providerInfo, ok := getProviderInfo(providerConfig.Name)
		if !ok {
			return nil, fmt.Errorf("provider not found: %s", providerConfig.Name)
		}
		providersInfo = append(providersInfo, providerInfo)
		provider, err := providerInfo.New(providerConfig.ClientKey, providerConfig.Secret, providerConfig.redirectURI.String(), providerConfig.Custom, providerConfig.Scopes...)
		if err != nil {
			return nil, fmt.Errorf("failed to create provider %s: %w", providerConfig.Name, err)
		}
		goth.UseProviders(provider)
	}
	return providersInfo, nil
}
