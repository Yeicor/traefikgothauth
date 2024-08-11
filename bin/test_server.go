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
	cfg.LogLevel = "trace"
	cfg.ProviderName = "twitch"
	cfg.ProviderCallback = "http://localhost:8080/__goth/"
	cfg.ProviderParams = map[string]interface{}{
		"clientKey": os.Getenv("TWITCH_CLIENT_KEY"),
		"secret":    os.Getenv("TWITCH_SECRET"),
	}
	cfg.CookieSecret = "secret-for-testing-only"
	cfg.ClaimsPrefix = "X-User-"
	cfg.Authorize = &traefikgothauth.AuthorizeConfig{
		Regexes: map[string]string{
			"user-id": "000000000", // Only allow user 000000000
		},
	}

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
