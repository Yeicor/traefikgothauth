package traefikgothauth_test

import (
	"context"
	"github.com/Yeicor/traefikgothauth"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := traefikgothauth.CreateConfig()
	cfg.ProviderName = "twitch"
	cfg.ProviderCallback = "http://localhost:8080/__goth/twitch/"
	cfg.ProviderParams = map[string]interface{}{
		"clientKey": os.Getenv("TWITCH_CLIENT_KEY"),
		"secret":    os.Getenv("TWITCH_SECRET"),
	}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikgothauth.New(ctx, next, cfg, "oidc")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	println(recorder.Code)
	// println(recorder.Header().Get("Location"))
}
