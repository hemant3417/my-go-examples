package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	fmt.Println(string(bs))
}
