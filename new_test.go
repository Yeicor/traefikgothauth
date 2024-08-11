package traefikgothauth_test

import (
	"context"
	"github.com/Yeicor/traefikgothauth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := traefikgothauth.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikgothauth.New(ctx, next, cfg, "oidc-plugin")
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
