package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, p.last, " and age is ", p.age)
}

func main() {
	p1 := person{
		first: "Henk",
		last:  "rovi",
		age:   45,
	}
	p1.speak()
}
