package main

import "testing"

func TestValidateImportCycle(t *testing.T) {
	cyclic := map[string][]string{
		"pkg1": {"pkg2"},
		"pkg2": {"pkg1"},
	}
	if !ValidateImportCycle(cyclic, "pkg1") {
		t.Error("expected cycle to be detected")
	}

	acyclic := map[string][]string{
		"cmd":     {"service"},
		"service": {"repository"},
	}
	if ValidateImportCycle(acyclic, "cmd") {
		t.Error("acyclic graph should not detect cycle")
	}
}
