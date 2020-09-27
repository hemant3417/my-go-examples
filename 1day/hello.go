package main

import "fmt"

var t = 89
var b int
var c string = `Hemanth says
               "Happy"`

type rambo int

var h rambo

func main() {
	n, err := fmt.Println("Hello Hi")
	fmt.Println(n)
	fmt.Println(err)

	n, _ = fmt.Println(Hello())
	x := 65
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%v\t%b\t%x\t%#x\n", x, x, x, x)
	s := fmt.Sprintf("%b\t%x\t%#x\t", x, x, x)
	fmt.Println(s)

	h = 75
	fmt.Println(h)
	fmt.Printf("%T\n", h)

	x = int(h)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(x)
	y := "Hemant"
	fmt.Println(y)
	fmt.Println(b)
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}

func Hello() string {
	return "Hello, world."
}
