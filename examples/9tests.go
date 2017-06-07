package main

import (
	"testing"
)

// In a <same_name>_test.go //HL
func TestSomething(t *testing.T) {
	result := 5 * 5
	if result != 10 {
		t.Error("Up to no good")
	}
}