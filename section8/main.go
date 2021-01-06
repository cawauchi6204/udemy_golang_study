package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
	// deferを使うことでTestDeferが終わった後にdefer分の処理を実行する
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

// defer文が2件以上登録されている場合、したから実行される

func main() {
	TestDefer()
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
