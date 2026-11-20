package main

import "testing"

func TestExtractLinks(t *testing.T) {
	htmlDoc := `
	<div>
		<a href="catalog">Catalog</a>
		<a href="/deals">Deals</a>
		<a href="https://external.com/page">External</a>
	</div>
	`
	links, err := ExtractLinks("https://shop.com/store/", htmlDoc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://shop.com/store/catalog",
		"https://shop.com/deals",
		"https://external.com/page",
	}

	if len(links) != len(expected) {
		t.Fatalf("expected %d links, got %d", len(expected), len(links))
	}
	for i, l := range links {
		if l != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], l)
		}
	}
}
