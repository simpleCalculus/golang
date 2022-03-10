package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var f1 int = -1
	var f2 int = 1
	return func() int {
		f := f1 + f2
		f1 = f2
		f2 = f
//		fmt.Printf("f1 = %d, f2 =%d\n", f1, f2)
		return f2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
