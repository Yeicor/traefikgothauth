package main

import (
	"context"
	"github.com/Yeicor/traefikgothauth"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := traefikgothauth.CreateConfig()
	cfg.Providers = []*traefikgothauth.ProviderConfig{
		{
			Name:        "twitch",
			ClientKey:   os.Getenv("TWITCH_CLIENT_KEY"),
			Secret:      os.Getenv("TWITCH_SECRET"),
			RedirectURI: "http://localhost:8080/__goth/twitch/",
		},
	}
	cfg.CookieSecret = "secret-for-testing-only"
	cfg.LogLevel = "trace"

	handler, err := traefikgothauth.New(context.Background(), http.HandlerFunc(backend), cfg, "oidc-plugin")
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}

func backend(rw http.ResponseWriter, req *http.Request) {
	// Dump the full request as the body of the response
	err := req.Write(rw)
	if err != nil {
		log.Println("Error writing response", "error", err)
	}
}
