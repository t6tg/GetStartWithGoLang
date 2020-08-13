package main

import "testing"

func TestSum(t *testing.T) {
	value := sum(2, 3)
	value2 := sum(5, 4)
	if value != 5 {
		t.Errorf("sum(2,3) must be 5 not %d", value)
	}
	if value2 != 9 {
		t.Errorf("sum(5,4) must be 9 not %d", value2)
	}
}
