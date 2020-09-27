package main

import (
	"fmt"
	"sort"
)

type Employ struct {
	Name string
	Age  int
}

type ByAge []Employ

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	e1 := Employ{
		"Roger",
		34,
	}

	e2 := Employ{
		"amber",
		30,
	}

	e3 := Employ{
		"bob",
		36,
	}

	ez := []Employ{e1, e2, e3}

	sort.Sort(ByAge(ez))

	fmt.Println(ez)

}
