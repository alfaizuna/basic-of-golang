package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

type dummySpeaker struct{}

func (d dummySpeaker) Speak() string {
	return "test speaker"
}

func TestSpeakerInterface(t *testing.T) {
	// Verifikasi dummy speaker bisa masuk ke Broadcast
	if got := Broadcast(dummySpeaker{}); got != "test speaker" {
		t.Errorf("expected 'test speaker', got %q", got)
	}

	bot := Robot{Model: "T-800"}
	// Verifikasi Robot memenuhi interface Speaker
	var s Speaker = bot
	if s.Speak() != "Beep boop! I am robot T-800" {
		t.Errorf("unexpected Speak result: %q", s.Speak())
	}
	if got := Broadcast(bot); got != "Beep boop! I am robot T-800" {
		t.Errorf("Broadcast failed: %q", got)
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
	want := "Robot speaks: Beep boop! I am robot RX-78"

	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q, got:\n%s", want, got)
	}
}
