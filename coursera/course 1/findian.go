package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func iAN(str string) string {
	tlstr := strings.ToLower(str)
	if !strings.HasPrefix(tlstr, "i") {
		return "Not Found!"
	}

	if !strings.HasSuffix(tlstr, "n") {
		return "Not Found!"
	}

	if !strings.Contains(tlstr, "a") {
		return "Not Found!"
	}

	return "Found!"
}

func main() {
	// var line string
	// fmt.Scan(&line)
	// fmt.Println(iAN(line))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name := scanner.Text()
		fmt.Println(iAN(name))
		break
	}
}
