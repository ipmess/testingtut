package main

import "testing"

// Test function for Add
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	// Check if the result is as expected
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}