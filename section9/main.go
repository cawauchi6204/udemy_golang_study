package main

import "fmt"

// slice

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for i := range sl {
	// 	fmt.Println(i)
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}
