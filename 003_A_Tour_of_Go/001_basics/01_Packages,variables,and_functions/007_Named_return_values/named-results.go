package main

import "fmt"

func split(sum int) (x, y int) {
	// 関数の宣言で戻り値に名前をつけてる
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
