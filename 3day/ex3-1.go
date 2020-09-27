package main

import "fmt"

func main() {
	x := [5]int{2, 8, 3, 1, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])

	y = append(y, 52)
	y = append(y, 53, 54, 55)

	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)

	y = y[:10]
	fmt.Println(y)
	y = append(y[:3], y[6:]...)
	fmt.Println(y)

	fmt.Printf("%T\n", y)
	usa := make([]string, 50, 50)
	usa = []string{" Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado",
		" Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois",
		" Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts",
		" Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada",
		" New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota",
		" Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota",
		" Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin",
		" Wyoming"}

	fmt.Println(len(usa))
	fmt.Println(cap(usa))
	for i := 0; i < len(usa); i++ {
		fmt.Println(i, usa[i])
	}

	xs := []string{"james", "bond", "shaken,not stirred"}
	ys := []string{"Miss", "moneypenny", "hello,james"}
	rs := [][]string{xs, ys}
	fmt.Println(rs)

	for i, xs1 := range rs {
		fmt.Println(i)
		for _, v := range xs1 {
			fmt.Println(v)
		}
	}

	hw := map[string][]string{
		"james":      {"hero", "women", "martini"},
		"moneypenny": {"jeame", "cs", "love"},
		"q":          {"cars", "guns", "missiles"},
	}

	fmt.Println(hw["james"])

	for k, v := range hw {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("Index %v\t and value %v\n", i, val)
		}
	}

	hw["joker"] = []string{"funny", "villian", "dangerous"}

	delete(hw, "q")

	for _, ek := range hw {
		fmt.Println(ek)
	}

}
