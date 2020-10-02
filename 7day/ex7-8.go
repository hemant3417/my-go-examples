package main

import (
	"fmt"
)

type custError struct {
	info string
}

func (c custError) Error() string {
	return fmt.Sprintf("Custom Error %v", c.info)
}

func main() {
	c1 := custError{
		"buggy",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran e - \n ", e.(custError).info)
}
