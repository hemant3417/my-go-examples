package main

import "fmt"

const (
	g     = 2017 + iota
	h     = 2017 + iota
	i     = 2017 + iota
	j     = 2017 + iota
	k int = 34
	l     = 56
)

func main() {
	x := 24
	fmt.Printf("x is %d in binary %b\n", x, x)
	fmt.Printf("x is %d in binary %#x\n", x, x)

	a := (7 == 21)
	b := (7 != 21)
	c := (7 <= 21)
	d := (7 >= 21)
	e := (7 < 21)
	f := (7 > 21)
	fmt.Println(a, b, c, d, e, f)

	fmt.Println(k)
	fmt.Println(l)

	s := 8
	t := s << 1
	fmt.Println("\n")
	fmt.Printf("%d\t%b\t%#x\n", s, s, s)
	fmt.Printf("Value of t is %d and binary is %b and hex is %#x\n", t, t, t)

	u := `Here us a raw 
		 string literal`
	fmt.Println(u)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

}
