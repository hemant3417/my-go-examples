package main

import (
	"fmt"
	"sort"
)

func main() {

	ii := []int{34, 12, 45, 23, 11}
	si := []string{"jock", "log", "aple", "som"}
	sort.Ints(ii)
	sort.Strings(si)
	fmt.Println(ii)
	fmt.Println(si)
}
