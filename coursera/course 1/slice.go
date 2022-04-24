package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type myNums []int64

func (this myNums) Len() int {
	return len(this)
}

func (this myNums) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this myNums) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortAndPrint(nums *[]int64) {
	mn := myNums(*nums)
	sort.Sort(mn)
	for _, v := range mn {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	nums := make([]int64, 0, 3)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter integers or X to exit the program")
	for {
		num, _ := reader.ReadString('\n')
		num = strings.Replace(num, "\n", "", -1)
		if num == "X" {
			break
		}
		intNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Print(err)
			break
		}
		nums = append(nums, intNum)
		sortAndPrint(&nums)
	}
}
