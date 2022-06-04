package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func defineMethod(a Animal, method string) string {
	switch method {
	case "eat":
		return a.Eat()
	case "move":
		return a.Move()
	case "speak":
		return a.Speak()
	default:
		return "not correct method"
	}
}

func activity() {
	fmt.Println(">")

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		var animal string
		var command string

		fmt.Println("enter name of an animal and name of the information requested about the animal")
		fmt.Scan(&animal, &command)

		switch animal {
		case "cow":
			fmt.Println(defineMethod(cow, command))
		case "bird":
			fmt.Println(defineMethod(bird, command))
		case "snake":
			fmt.Println(defineMethod(snake, command))
		default:
			fmt.Println("not correct input")
		}
	}
}

func main() {
	activity()
}
