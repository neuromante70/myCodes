package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := float64(1.0)
	y := float64(z)
	for {
		y -= (y*y - x) / (2.0 * y)
		if y*y-x <= 0.5 {
			return z
		}
		z++
	}
}

func main() {
	fmt.Println(sqrt(100))
}
