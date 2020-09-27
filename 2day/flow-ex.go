package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("The outer loop value is %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop value is %d \n", j)
		}
	}

	x := 0
	for {
		x++
		if x == 2 {
			continue
		}
		if x > 10 {
			break
		}
		fmt.Println(x)
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

	if x := 42; x != 2 {
		fmt.Println("Here comes")
	}

	switch {
	case false:
		fmt.Println("No")
	case 2 == 2:
		fmt.Println("Yes 2==2")
		fallthrough
	case 3 == 3:
		fmt.Println("Yes 3==3")
		fallthrough
	default:
		fmt.Println("This is default")
	}

	n := "Bond"
	switch n {
	case "MoneyPenny", "K":
		fmt.Println("MoneyPenny")
	case "Bond":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("Q")
	}

}
