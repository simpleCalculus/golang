package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const EPS float64 = 1e-15
	z := 1.0
	for t := 2.0; math.Abs(z-t) > EPS; {
		t = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("Sqrt: %f\n", Sqrt(802))
	fmt.Printf("math.Sqrt: %f\n", math.Sqrt(802))
}
