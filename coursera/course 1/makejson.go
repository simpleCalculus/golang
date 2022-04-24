package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string
	fmt.Scan(&name)
	fmt.Scan(&addr)

	myMap := make(map[string]string)
	myMap[name] = addr

	barr, _ := json.Marshal(myMap)
	fmt.Print(barr, string(barr))
}
