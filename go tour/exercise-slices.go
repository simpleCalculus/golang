package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	myPic := make([][]uint8, dy)
	for i := range myPic {
		myPic[i] = make([]uint8, dx)
	}
	for i := range myPic {
		for j := range myPic[i] {
			myPic[i][j] = uint8((i + j))
		}
	}

	return myPic
}

func main() {
	pic.Show(Pic)
}
