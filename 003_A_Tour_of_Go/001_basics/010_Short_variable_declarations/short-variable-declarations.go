package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// var宣言の代わりに`:=`で暗黙的な型宣言ができる！
	// ただ、関数の外ではつかえない！！
	k := 3
	c, python, java := true, false, "no!!"

	fmt.Println(i, j, k, c, python, java)
}
