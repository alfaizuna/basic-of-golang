package main

import "testing"

func TestQueueJobs(t *testing.T) {
	jobs := []string{"email", "sms", "push"}
	ch, capacity := QueueJobs(jobs)

	if capacity != 3 {
		t.Fatalf("expected capacity 3, got %d", capacity)
	}

	for _, expected := range jobs {
		select {
		case got := <-ch:
			if got != expected {
				t.Errorf("expected job %q, got %q", expected, got)
			}
		default:
			t.Fatalf("channel was empty earlier than expected")
		}
	}
}
