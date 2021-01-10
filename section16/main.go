package main

import (
	"log"
	"os"
)

//time

func main() {
	// log.SetOutput(os.Stdout)
	// スタンダードアウト

	/*
		log.Print("Log\n")
		log.Println("Log2")
		log.Printf("Log%d\n", 3)
	*/

	/*
		log.Fatal("Log\n")
		log.Fatalln("Log2")
		log.Fatalf("Log%d\n", 3)
	*/

	f, err := os.Create("test.log")
	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println("ファイルに書き込み")

}
