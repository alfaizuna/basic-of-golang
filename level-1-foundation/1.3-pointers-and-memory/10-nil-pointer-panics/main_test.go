package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	prof := createProfile("Backend Developer")
	if prof == nil {
		t.Fatal("createProfile mengembalikan nil, seharusnya pointer ke struct Profile yang valid")
	}
	if prof.Bio != "Backend Developer" {
		t.Errorf("createProfile: bio ingin %q, dapat %q", "Backend Developer", prof.Bio)
	}
}

func TestLoadConfig(t *testing.T) {
	cfg := loadConfig("PaymentService", "127.0.0.1", 9000)
	if cfg == nil {
		t.Fatal("loadConfig mengembalikan nil, seharusnya pointer ke struct AppConfig yang valid")
	}
	if cfg.AppName != "PaymentService" {
		t.Errorf("loadConfig: AppName ingin %q, dapat %q", "PaymentService", cfg.AppName)
	}
	if cfg.Server == nil {
		t.Fatal("loadConfig: cfg.Server masih nil, seharusnya dialokasikan dengan pointer ke ServerConfig")
	}
	if cfg.Server.Host != "127.0.0.1" {
		t.Errorf("loadConfig: Server.Host ingin %q, dapat %q", "127.0.0.1", cfg.Server.Host)
	}
	if cfg.Server.Port != 9000 {
		t.Errorf("loadConfig: Server.Port ingin %d, dapat %d", 9000, cfg.Server.Port)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

	main()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	wantProfile := "Profile: Software Engineer"
	wantConfig := "Config: WebApp on localhost:8080"

	if !strings.Contains(got, wantProfile) {
		t.Errorf("output tidak mengandung %q, dapat:\n%s", wantProfile, got)
	}
	if !strings.Contains(got, wantConfig) {
		t.Errorf("output tidak mengandung %q, dapat:\n%s", wantConfig, got)
	}
}
