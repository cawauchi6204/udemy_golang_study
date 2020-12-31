package main

import "fmt"

const Pi = 3.14

const (
	URL      = "https://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 頭文字を大文字にすることで他のファイルから取得できる
//定数はグローバルに書くことが多い

func main() {
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
}
