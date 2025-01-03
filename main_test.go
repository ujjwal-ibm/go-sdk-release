package main

import "testing"

func TestMain(t *testing.T) {
	result := "Welcome to the Go SDK!"
	if result != "Welcome to the Go SDK!" {
		t.Errorf("Unexpected output: %s", result)
	}
}
