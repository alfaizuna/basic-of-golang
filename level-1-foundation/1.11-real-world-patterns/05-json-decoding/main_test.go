package main

import "testing"

func TestParseConfig(t *testing.T) {
	raw := []byte(`{"port": 9000, "host": "127.0.0.1", "debug": false}`)
	cfg, err := ParseConfig(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != 9000 || cfg.Host != "127.0.0.1" || cfg.Debug != false {
		t.Errorf("fields parsed incorrectly: %+v", cfg)
	}

	_, err = ParseConfig([]byte("invalid json"))
	if err == nil {
		t.Error("expected error on malformed json")
	}
}
