package main

import "fmt"

type person1 struct {
	first string
	last  string
}

type sAgent1 struct {
	person1
	ltk bool
}

func (s sAgent1) speak() {
	fmt.Println("I am ", s.first, s.last, s.ltk)
}

func (s person1) speak() {
	fmt.Println("HI there ", s.first, s.last)
}

type human interface {
	speak()
}

func bar2(h human) {
	switch h.(type) {
	case person1:
		fmt.Println("I am passed into Barrr ", h.(person1).first)
	case sAgent1:
		fmt.Println("I am passed into Barrr ", h.(sAgent1).first)

	}
	fmt.Println("I am passed into bar", h)
}

type hotdog int

func main() {

	sa1 := sAgent1{
		person1: person1{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := sAgent1{
		person1: person1{
			first: "MOney",
			last:  "Penny",
		},
		ltk: false,
	}

	p21 := person1{
		first: "Jamie111",
		last:  "fox",
	}

	sa1.speak()
	sa2.speak()

	fmt.Println(sa1)

	bar2(sa1)
	bar2(sa2)
	bar2(p21)

	defer foo()
	bar("Q")
	s1 := woo("ram")
	fmt.Println(s1)
	xr, yr := moo("Jamie", "fox")
	fmt.Println(xr, yr)

	f := func(x int) {
		fmt.Println("Hi.... ", x)
	}
	f(1990)

	sum := calcTotal(2, 3, 4, 56, 45)
	fmt.Println(sum)

	xi := []int{34, 45, 67, 89}
	total := calcTotal(xi...)
	fmt.Println(total)

	var x hotdog = 42
	fmt.Printf("%T\n", x)
	var y int = int(x)
	fmt.Printf("%T\n", y)

	fmt.Println("End of main")

	func(x int) {
		fmt.Println(x)
	}(42)

	f2 := func() {
		fmt.Println("Fun expressions")
	}

	f2()

	cr := fav()
	fmt.Printf("%T\n", cr)

	crv := cr()
	fmt.Println(crv)

}

func fav() func() int {
	return func() int {
		return 4223
	}
}

func foo() {
	fmt.Println("In foo")
}

func bar(s string) {
	fmt.Println("Hello ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello ", s)
}

func moo(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `says, "Hello"`)
	b := false
	return a, b
}

func calcTotal(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
