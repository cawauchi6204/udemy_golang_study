package main

import "fmt"

// map

func main() {
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}

	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])

	_, ok := m4[1]
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println("ok")
	}

	delete(m4, 2)
	fmt.Println(m4)
	// fmt.Println(s, ok)
}
