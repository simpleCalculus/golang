package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func convertToStruct(line string) Person {
	strs := strings.Split(line, " ")
	return Person{strs[0], strs[1]}
}

func main() {
	osReader := bufio.NewReader(os.Stdin)
	path, err := osReader.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}
	path = strings.Replace(path, "\n", "", -1)

	src, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		if cerr := src.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	data := make([]Person, 0)

	fileReader := bufio.NewReader(src)
	for {
		line, err := fileReader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err == io.EOF {
			data = append(data, convertToStruct(line))
			break
		}
		if err != nil {
			log.Print(err)
			return
		}
		data = append(data, convertToStruct(line))
	}

	for _, person := range data {
		fmt.Println(person.fname, person.lname)
	}
}
