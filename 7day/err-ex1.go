package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer f.Close()
	r := strings.NewReader("James bond")
	io.Copy(f, r)
}
