package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSONPayload(t *testing.T) {
	var receivedBody string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		receivedBody = string(b)
		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()

	client := ts.Client()
	status, err := SendJSONPayload(client, ts.URL, []byte(`{"id": 1}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 201 || receivedBody != `{"id": 1}` {
		t.Errorf("got status=%d, body=%q", status, receivedBody)
	}
}
