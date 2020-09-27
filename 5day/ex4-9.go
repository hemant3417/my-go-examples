package main

import "fmt"

type persona struct {
	first string
	last  string
}

func changeme(p *persona) {
	p.first = "nyca"
	(*p).last = "ti"
}

func main() {

	pi := persona{
		"Jame",
		"lawson",
	}

	fmt.Println(pi.first, pi.last)

	changeme(&pi)

	fmt.Println(pi.first, pi.last)

}
