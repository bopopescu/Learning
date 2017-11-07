package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
