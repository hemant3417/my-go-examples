package main

import "testing"

func TestMySum(t *testing.T) {
	s := MySum(1, 2)
	if s != 3 {
		t.Error("Expected 3 got ", s)
	}
}
