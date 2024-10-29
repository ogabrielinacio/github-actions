package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)
	if total != 30 {
		t.Errorf("Invalid Result. Result: %d; expected: %d", total, 30)
	}
}
