package main

import (
	"testing"
)

func TestMySum1(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 3, 5}, 9},
		test{[]int{1, 2}, 3},
		test{[]int{11, 12}, 23},
	}

	for _, v := range tests {
		x := mySum1(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, " got ", x)
		}
	}
}
