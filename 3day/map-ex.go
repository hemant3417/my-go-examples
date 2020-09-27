package main

import "fmt"

func main() {
	m := map[string]int{
		"Jamee": 40,
		"Tracy": 28,
	}

	fmt.Println(m["Jamee"])

	if v, ok := m["Jamee"]; ok {
		fmt.Println("Jamee Present : ", v)
	}

	m["Tod"] = 45

	delete(m, "Jamee")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
