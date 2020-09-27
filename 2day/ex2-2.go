package main

import "fmt"

func main() {
	x := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(x)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
		x++
	}

	yob := 1987
	for yob <= 2020 {
		fmt.Println(yob)
		yob++
	}

	dy := 1987
	for {
		fmt.Println(dy)
		dy++
		if dy > 2020 {
			break
		}
	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("when %v is divided by 4 then modulus is %v\n", i, i%4)
	}
}
