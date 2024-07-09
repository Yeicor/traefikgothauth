package traefikgothauth

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/sethvargo/go-envconfig"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// EnvPrefix is the prefix for the environment variables. Any configuration can be set via environment variables.
const EnvPrefix = "TRAEFIKGOTHAUTH_"

// Config configures the Goth Auth plugin.
type Config struct {
	// LogLevel is the log level (trace, debug, info, warn, error, off).
	LogLevel string

	// ProviderName is the configured provider (see providers.go).
	ProviderName string
	// ProviderCallback is the callback URL for the provider (must be the same as in the provider's configuration).
	ProviderCallback string
	providerCallback *url.URL
	// ProviderParams are the provider-specific parameters (see the docs from Goth and the provider).
	ProviderParams map[string]interface{}

	// CookieSecret is the secret used to sign the cookie.
	CookieSecret string
	// CookieOptions are the cookie options.
	CookieOptions *sessions.Options

	// LogoutURI (optional) is the URI to logout from the provider. Defaults to no logout (manually remove the cookie).
	LogoutURI string
	logoutURI *url.URL

	// Authorize configures the checks to ensure the user is authorized. By default, any authenticated user is authorized.
	Authorize *AuthorizeConfig

	// ClaimsPrefix is the prefix for the claims to be published as headers. Defaults to not publishing any claims.
	ClaimsPrefix string
}

// AuthorizeConfig configures the checks to ensure the user is authorized.
type AuthorizeConfig struct {
	// Regexes are the regular expressions to match the user claims.
	Regexes map[string]string
	regexes map[string]*regexp.Regexp
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	providerParams := map[string]interface{}{}
	for _, s := range os.Environ() {
		if strings.HasPrefix(s, EnvPrefix+"PROVIDER_PARAMS_") {
			parts := strings.SplitN(s, "=", 2)
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimPrefix(parts[0], EnvPrefix+"PROVIDER_PARAMS_")
			providerParams[key] = parts[1]
		}
	}
	return &Config{
		ProviderParams: providerParams,
		CookieOptions:  &sessions.Options{HttpOnly: true, Path: "/", MaxAge: 60 * 60},
		ClaimsPrefix:   "__NO__",
		LogLevel:       "info",
	}
}

func (c *Config) setup() (*ProviderInfo, error) {
	// FIXME(yaegi): Extend config from environment variables
	err := envconfig.ProcessWith(context.Background(), &envconfig.Config{
		Target:   c,
		Lookuper: envconfig.PrefixLookuper(EnvPrefix, envconfig.OsLookuper()),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process environment variables: %w", err)
	}

	// Set up the loggers
	var ok bool
	logLevelCurrent, ok = logTextLevel[strings.ToUpper(c.LogLevel)]
	if !ok {
		loge("Invalid log level", "level", c.LogLevel)
		logLevelCurrent = logLevelInfo
	}

	// Prepare the cookie store
	// TODO: Can this global store cause conflicts between multiple plugin instances?
	//  If this is a problem a possible fix is to rewrite the (small) gothic package
	//  The same happens for goth.UseProviders unless we append a prefix to each one...
	//  The same happens for logLevelCurrent...
	gothic.Store = sessions.NewCookieStore([]byte(c.CookieSecret))
	gothic.Store.(*sessions.CookieStore).Options = c.CookieOptions

	// Parse and validate the inputs
	if c.LogoutURI != "" {
		c.logoutURI, err = url.Parse(c.LogoutURI)
		if err != nil {
			return nil, fmt.Errorf("failed to parse default logout URI: %w", err)
		}
	}
	c.providerCallback, err = url.Parse(c.ProviderCallback)
	if err != nil {
		return nil, fmt.Errorf("failed to parse provider callback: %w", err)
	}
	if c.Authorize != nil {
		c.Authorize.regexes = make(map[string]*regexp.Regexp, len(c.Authorize.Regexes))
		for k, v := range c.Authorize.Regexes {
			c.Authorize.regexes[k], err = regexp.Compile(v)
			if err != nil {
				return nil, fmt.Errorf("failed to compile regex: %w", err)
			}
		}
	}

	// Create the auth provider and set it up
	providerInfo, ok := getProviderInfo(c.ProviderName)
	if !ok {
		return nil, fmt.Errorf("provider not found: %s", c.ProviderName)
	}
	provider, err := providerInfo.New(c.ProviderCallback, c.ProviderParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}
	goth.UseProviders(provider)

	return providerInfo, nil
}
