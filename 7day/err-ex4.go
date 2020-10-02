package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("Error occured whil sqrt")
		return 0, fmt.Errorf("Error finding sqrt of %v ", f)
	}
	return 42, nil
}
