package traefikgothauth

import (
	"github.com/markbates/goth"
)

//go:generate go run -v bin/providers_gen.go

func getProviderInfo(name string) (*ProviderInfo, bool) {
	for _, p := range allProviders {
		if p.Name == name {
			return p, true
		}
	}
	return nil, false
}

// ProviderInfo contains static metadata for a provider.
type ProviderInfo struct {
	Name string
	New  func(callback string, custom map[string]interface{}) (goth.Provider, error)
}
