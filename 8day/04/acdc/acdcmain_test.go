package acdc

import (
	"fmt"
	"testing"
)

func ExampleSum() {
	fmt.Println(Sum(1, 2))
	// Output:
	// 3
}

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
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, " got ", x)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}
