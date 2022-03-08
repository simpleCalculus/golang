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

/*
 If you are interested in the details of the algorithm, the z² − x above is 
 how far away z² is from where it needs to be (x), and the division by 2z is 
 the derivative of z², to scale how much we adjust z by how quickly z² is 
 changing. This general approach is called Newton's method. It works well 
 for many functions but especially well for square root.
*/
