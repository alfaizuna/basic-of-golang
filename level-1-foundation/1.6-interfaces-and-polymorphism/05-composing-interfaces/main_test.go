package main

import "testing"

func TestTransceiverImplementation(t *testing.T) {
	var _ Transceiver = (*RadioDevice)(nil)

	radio := &RadioDevice{}
	res, err := BroadcastPing(radio, "Test Message")
	if err != nil {
		t.Fatalf("BroadcastPing returned unexpected error: %v", err)
	}
	if res != "Test Message" {
		t.Errorf("expected %q, got %q", "Test Message", res)
	}
}
