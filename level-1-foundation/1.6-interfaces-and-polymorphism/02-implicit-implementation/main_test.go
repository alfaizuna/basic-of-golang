package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestImplicitImplementation(t *testing.T) {
	email := EmailSender{Email: "admin@test.com"}
	sms := SMSSender{Phone: "089999"}

	// Verifikasi kedua tipe dapat di-assign ke Notifier
	var _ Notifier = email
	var _ Notifier = sms

	list := []Notifier{email, sms}
	alerts := BroadcastAlert(list, "Disk full")

	if len(alerts) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(alerts))
	}
	if alerts[0] != "[EMAIL to admin@test.com] Disk full" {
		t.Errorf("unexpected email alert: %q", alerts[0])
	}
	if alerts[1] != "[SMS to 089999] Disk full" {
		t.Errorf("unexpected sms alert: %q", alerts[1])
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	got := strings.TrimSpace(buf.String())
	want1 := "Notification 1: [EMAIL to user@example.com] Server down!"
	want2 := "Notification 2: [SMS to +62812345678] Server down!"

	if !strings.Contains(got, want1) {
		t.Errorf("output does not contain %q, got:\n%s", want1, got)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("output does not contain %q, got:\n%s", want2, got)
	}
}
