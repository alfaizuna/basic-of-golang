package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSumHandler(t *testing.T) {
	body := []byte(`{"a": 12, "b": 8}`)
	req := httptest.NewRequest(http.MethodPost, "/api/sum", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	SumHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var res SumResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if res.Sum != 20 {
		t.Errorf("expected sum 20, got %d", res.Sum)
	}
}
