package main

import "fmt"

// `変数名or関数名 型名`のように記述する！
// xとyの型が同じなので下記のように省略できる！
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
