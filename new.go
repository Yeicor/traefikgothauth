package traefikgothauth

import (
	"context"
	"github.com/gorilla/sessions"
	"net/http"
)

// Plugin is the Traefik Goth Auth plugin.
type Plugin struct {
	next          http.Handler
	config        *Config
	providersInfo []*ProviderInfo
	redirectStore sessions.Store
}

// New created a New Plugin plugin.
//
//goland:noinspection GoUnusedParameter
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	providers, err := config.setup()
	if err != nil {
		return nil, err
	}
	redirectStore := sessions.NewCookieStore([]byte(config.CookieSecret))
	*redirectStore.Options = *config.CookieOptions // Copy
	redirectStore.Options.MaxAge = 0               // Session only
	return &Plugin{
		next:          next,
		config:        config,
		providersInfo: providers,
		redirectStore: redirectStore,
	}, nil
}
