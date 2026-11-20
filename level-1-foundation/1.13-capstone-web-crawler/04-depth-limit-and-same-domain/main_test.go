package main

import "testing"

func TestIsAllowed(t *testing.T) {
	base := "store.com"

	// Valid same domain & within depth
	if !IsAllowed("https://store.com/item", base, 1, 3) {
		t.Error("expected true for valid link")
	}

	// Exceeded depth
	if IsAllowed("https://store.com/item", base, 4, 3) {
		t.Error("expected false when depth exceeds maxDepth")
	}

	// External domain
	if IsAllowed("https://google.com", base, 1, 3) {
		t.Error("expected false for external domain")
	}
}
