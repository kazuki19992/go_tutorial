package main

import "fmt"

func main() {
	sum := 1

	// セミコロンの省略もできる！
	// つまり、C言語とかの`while`は`for`をつかう！
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
