package main

import "fmt"

const Pi = 3.14 // 定数
// 定数は := で宣言できない！

func main() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
