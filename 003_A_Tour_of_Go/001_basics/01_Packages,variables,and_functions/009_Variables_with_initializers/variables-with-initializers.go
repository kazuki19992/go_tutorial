package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 初期化子が与えられている場合、型は省略できる！
	var c, python, java = true, false, "no!!!"
	fmt.Println(i, j, c, python, java)
}
