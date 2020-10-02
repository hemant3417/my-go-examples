package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error occured")
	}

	log.SetOutput(f)

	defer f.Close()

	f2, err := os.Open("no-such-file.txt")
	if err != nil {
		//fmt.Println("Error happened-> : ", err)
		//log.Println("Error happened-> : ", err)
		//log.Fatalln(err)
		//log.Panic(err)
		panic(err)
	}

	defer f2.Close()

	fmt.Println("Check the log.txt file for error logs")

}

func foo() {
	fmt.Println("In foo testing log fatalln")
}
