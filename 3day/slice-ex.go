package main

import (
	"fmt"
)

func main() {
	x := []int{3, 6, 1, 56, 45}
	fmt.Println(x)
	x = append(x, 43, 57, 23)

	y := []int{21, 65, 723}
	x = append(x, y...)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	x = append(x[:2], x[4:]...)
	fmt.Println("-->>", x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	z := make([]int, 10, 100)
	z = append(x, 45)
	fmt.Println(len(z))

	jp := []string{"axe", "boo", "caa", "duuu"}
	xp := []string{"voo", "xoo", "yo", "oz"}
	mb := [][]string{jp, xp}
	fmt.Println(mb)
}
