package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// 初期値を与えないと以下のようになる！
	// - 数値型(int, frortなど)	：0
	// - bool型					：false
	// - string型				：""(空文字列)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
