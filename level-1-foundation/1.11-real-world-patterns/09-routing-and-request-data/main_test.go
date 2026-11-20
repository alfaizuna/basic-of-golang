package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserSearchHandler(t *testing.T) {
	// Test POST rejected
	postReq := httptest.NewRequest(http.MethodPost, "/users", nil)
	postRec := httptest.NewRecorder()
	UserSearchHandler(postRec, postReq)
	if postRec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405 on POST, got %d", postRec.Code)
	}

	// Test GET with query param
	getReq := httptest.NewRequest(http.MethodGet, "/users?q=golang", nil)
	getRec := httptest.NewRecorder()
	UserSearchHandler(getRec, getReq)
	if getRec.Code != 200 || getRec.Body.String() != "Query: golang" {
		t.Errorf("expected 'Query: golang', got %q", getRec.Body.String())
	}
}
