package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// Get the PATH environment variable
	pathString := os.Getenv("PATH")
	listOfPaths := strings.Split(pathString, string(filepath.ListSeparator))

	// Sort list if sorting flag provided
	toSort := toSort()
	if toSort {
		sort.Strings(listOfPaths)
	}

	// Print the paths
	printPaths(listOfPaths)
}

func toSort() bool {
	// Parse command-line flags
	sortFlag := flag.Bool("s", false, "sort the variables")
	sortFlagAlias := flag.Bool("sort", false, "sort the variables")
	flag.Parse()

	toSort := false
	if *sortFlag || *sortFlagAlias {
		toSort = true
	}

	return toSort
}

func printPaths(paths []string) {
	for _, p := range paths {
		fmt.Println(p)
	}
}
