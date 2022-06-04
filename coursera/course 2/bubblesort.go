package main

import "fmt"

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func BubbleSort(nums []int) {
	for i := range nums {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func main() {
	var nums []int
	fmt.Println("Enter 10 integer numbers")
	for i := 0; i < 10; i++ {
		var a int
		fmt.Scan(&a)
		nums = append(nums, a)
	}
	BubbleSort(nums)
	for _, v := range nums {
		fmt.Print(v, " ")
	}
}
