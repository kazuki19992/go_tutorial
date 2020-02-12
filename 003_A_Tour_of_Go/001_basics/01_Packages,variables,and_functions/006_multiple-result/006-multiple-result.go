package main

import "fmt"

// 関数名(引数) (戻り値の型(関数の型))
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
