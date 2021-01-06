package main

import "fmt"

// init
// mainルーティンよりも先に行われるのがinitルーティン
// 複数

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Main")
}
