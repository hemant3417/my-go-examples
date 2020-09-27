package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person2 struct {
	First string
	Last  string
}

func main() {

	p1 := Person2{
		"Jami",
		"lawson",
	}

	p2 := Person2{
		"morg",
		"fgh",
	}

	people := []Person2{p1, p2}
	fmt.Println(people)

	//bs, err := json.Marshal(people)

	/*
		if err != nil {
			fmt.Println("Error")
		}
	*/
	//fmt.Println(string(bs))

	errr := json.NewEncoder(os.Stdout).Encode(people)

	if errr != nil {
		fmt.Println("error encoding")
	}

}
