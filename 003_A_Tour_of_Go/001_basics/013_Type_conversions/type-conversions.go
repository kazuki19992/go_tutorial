package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// goでは明示的に型変換する必要がある！！
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
