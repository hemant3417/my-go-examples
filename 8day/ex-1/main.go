package main

import (
	"fmt"

	"github.com/my-examples/8day/ex-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	c1 := canine{
		"fredo", dog.Years(10),
	}
	fmt.Println(c1)

}
