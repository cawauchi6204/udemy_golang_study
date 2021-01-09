package main

import (
	"fmt"
	"math/rand"
	"time"
)

//time

func main() {
	// rand.Seed(256)

	// fmt.Println(rand.Float64())
	// fmt.Println(rand.Float64())
	// fmt.Println(rand.Float64())

	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
}
