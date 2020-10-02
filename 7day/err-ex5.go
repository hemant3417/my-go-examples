package main

import (
	"fmt"
	"log"
)

type northgateError struct {
	lat string
	lon string
	err error
}

func (n northgateError) Error() string {
	return fmt.Sprintf("Northgate error occured at %v %v %v", n.lat, n.lon, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nes := fmt.Errorf("Error occured finding sqrt of %v", f)
		return 0, northgateError{"43.2", "44.6", nes}
	}
	return 0, nil
}
