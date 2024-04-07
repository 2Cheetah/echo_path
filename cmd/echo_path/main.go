package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	pathString := os.Getenv("PATH")
	listOfPaths := strings.Split(pathString, ":")

	sortFlag := flag.Bool("s", false, "sort the variables")
	sortFlagAlias := flag.Bool("sort", false, "sort the variables")

	flag.Parse()

	if *sortFlag || *sortFlagAlias {
		sort.Strings(listOfPaths)
	}

	for _, s := range listOfPaths {
		fmt.Println(s)
	}
}
