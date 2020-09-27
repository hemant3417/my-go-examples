package main

import (
	"fmt"
	"runtime"
)

var a bool
var p int
var q float64
var r int8

const (
	j         = iota + 1
	k         = iota
	l         = iota
	g         = 34
	h float64 = 34.24
	i         = "James Bond"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println(a)
	a = true
	fmt.Println(a)
	x, y := 7, 7
	fmt.Println(x == y)

	p = 35
	q = 35.678
	fmt.Println(p)
	fmt.Println(q)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", q)
	r = -128
	fmt.Println(r)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	v := "Hello Playground"
	fmt.Println(v)
	fmt.Printf("%T\n", v)

	bs := []byte(v)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(v); i++ {
		fmt.Printf("%#U ", v[i])
	}

	fmt.Printf("\n")

	for i, v := range v {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", i)

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	m := 4
	fmt.Println(m)
	fmt.Printf("%d\t\t%b\n", m, m)

	n := m << 1
	fmt.Println(n)
	fmt.Printf("%d\t\t%b\n", n, n)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
