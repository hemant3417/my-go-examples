package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	awd bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		first: "Jammie",
		last:  "fox",
		fav:   []string{"chcoco", "vanilla"},
	}

	p2 := person{
		first: "Hillary",
		last:  "awank",
		fav:   []string{"mango", "raspberry"},
	}

	fmt.Println(p1.first, p1.last)

	for _, v := range p1.fav {
		fmt.Println(v)
	}

	fmt.Println(p2.first, p2.last)

	for i, v := range p2.fav {
		fmt.Println(i, v)
	}

	fmt.Println(p2)

	pe := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range pe {
		fmt.Println(k)
		fmt.Println(v.first, v.last)
		for i, val := range v.fav {
			fmt.Println(i, val)
		}
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		awd: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	e1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "roger",
		friends: map[string]int{
			"ranger": 35,
			"rocky":  45,
		},
		favDrinks: []string{"Cola", "Bud"},
	}

	fmt.Println(e1)

	for i, v := range e1.friends {
		fmt.Println(i, v)
	}

}
