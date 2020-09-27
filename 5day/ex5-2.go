package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

func main() {

	s := `[{"First":"Jami","Last":"lawson"},{"First":"morg","Last":"fgh"}]`

	bs := []byte(s)

	var people []Person2

	json.Unmarshal(bs, &people)

	fmt.Println(people)

}
