package main

import (
	"context"
	traefikoidc "github.com/Yeicor/traefik-oidc"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := traefikoidc.CreateConfig()
	cfg.ClientID = os.Getenv("OIDC_CLIENT_ID")
	cfg.ClientSecret = os.Getenv("OIDC_CLIENT_SECRET")
	cfg.ProviderURL = os.Getenv("OIDC_PROVIDER_URL")

	backend := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Dump the full request as the body of the response
		err := req.Write(rw)
		if err != nil {
			slog.Error("Error writing response", "error", err)
		}
	})

	handler, err := traefikoidc.New(context.Background(), backend, cfg, "oidc-plugin")
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
