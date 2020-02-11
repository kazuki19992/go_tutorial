package main

import "fmt"

// `変数名or関数名 型名`のように記述する！
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
