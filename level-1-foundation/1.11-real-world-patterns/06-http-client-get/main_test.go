package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchStatusCodeAndBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("mock server response"))
	}))
	defer ts.Close()

	status, body, err := FetchStatusCodeAndBody(ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 || body != "mock server response" {
		t.Errorf("got status=%d, body=%q", status, body)
	}
}
