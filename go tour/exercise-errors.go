package main

import (
	"fmt"
	"math"
	"errors"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0.0, errors.New(e.Error())
	}
	const EPS float64 = 1e-15
	z := 1.0
	for t := 2.0; math.Abs(z-t) > EPS; {
		t = z
		z -= (z*z - x) / (2 * z)
	}
	
	return z, nil
}

func Sqrt2(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	const EPS float64 = 1e-15
	z := 1.0
	for t := 2.0; math.Abs(z-t) > EPS; {
		t = z
		z -= (z*z - x) / (2 * z)
	}
	
	return z, 0
}

func main() {
	s, e := Sqrt(-2)
	if e == nil {
		fmt.Println(s)
	} else {
		fmt.Println(e)
	}
	sq, er := Sqrt2(-2)
	if er != 0 {
		fmt.Println(er.Error())
	} else {
		fmt.Println(sq)
	}
}
