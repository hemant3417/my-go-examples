package main

import "fmt"

func main() {

	a := 42

	fmt.Println(a)
	fmt.Println(&a) // & give us the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a

	fmt.Println(b)
	fmt.Println(*b) // * give us the value at address
	fmt.Println(*&a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)

	*b = 67
	fmt.Println(a)

	x := 0
	fmt.Println("x value ", x)
	fmt.Println("x adrees ", &x)
	foo(&x)

}

func foo(y *int) {
	fmt.Println("befor y address: ", y)
	fmt.Println("befor y value: ", *y)

	*y = 56
	fmt.Println("after y address: ", y)
	fmt.Println("after y value: ", *y)

}
