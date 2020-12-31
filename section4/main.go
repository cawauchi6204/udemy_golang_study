package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)

	x = 2
	// interfaceは全てと互換性を持つだけで組み合わせることができない
}
