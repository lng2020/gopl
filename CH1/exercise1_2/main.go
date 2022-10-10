package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result string
	for index, args := range os.Args[0:] {
		result = strconv.Itoa(index) + " " + args
		fmt.Println(result)
	}
}
