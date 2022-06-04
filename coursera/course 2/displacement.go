package main

import "fmt"

func GenDisplaceFn(acceleration, velocity, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		s := 0.5*acceleration*t*t + velocity*t + s0
		return s
	}
}

func main() {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&initialVelocity)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&initialDisplacement)
	fmt.Print("Enter time: ")
	fmt.Scan(&time)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(fn(3))
}
