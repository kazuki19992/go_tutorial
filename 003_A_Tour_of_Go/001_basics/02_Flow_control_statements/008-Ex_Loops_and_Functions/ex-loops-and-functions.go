// わからんちんちん

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for (x-z > 0.01) || (z-x > 0.1) {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
