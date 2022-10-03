package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, "回目の結果：", z)

		if math.Sqrt(x) == z {
			fmt.Println("OK!!!")
			break
		}
	}
	return z
}

func main() {
	Sqrt(5)
}
