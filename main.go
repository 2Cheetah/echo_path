package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pathString := os.Getenv("PATH")
	listOfPaths := strings.Split(pathString, ":")
	for _, s := range listOfPaths {
		fmt.Println(s)
	}
}
