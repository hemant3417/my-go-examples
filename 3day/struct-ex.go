package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   45,
		},
		ltk: true,
	}

	p2 := person{
		first: "Jamie",
		last:  "fox",
		age:   52,
	}

	p3 := struct {
		first string
		last  string
	}{
		first: "roj",
		last:  "rty",
	}
	fmt.Println(p3)
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2)

}
