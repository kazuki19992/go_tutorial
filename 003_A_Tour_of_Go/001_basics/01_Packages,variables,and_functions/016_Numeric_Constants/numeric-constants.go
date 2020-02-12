package main

import "fmt"

// intは64ビットの整数を保持できるが、それでは足りないときは
// constを活用していこう！！
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	Big = 1 << 100

	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
