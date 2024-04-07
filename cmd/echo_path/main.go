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
	// Parse command-line flags
	sortFlag := flag.Bool("s", false, "sort the variables")
	sortFlagAlias := flag.Bool("sort", false, "sort the variables")
	flag.Parse()

	// Get the PATH environment variable
	pathString := os.Getenv("PATH")
	listOfPaths := strings.Split(pathString, string(filepath.ListSeparator))

	// Sort the paths if the sort flag is provided
	if *sortFlag || *sortFlagAlias {
		sort.Strings(listOfPaths)
	}

	// Print the paths
	printPaths(listOfPaths)
}

func printPaths(paths []string) {
	for _, p := range paths {
		fmt.Println(p)
	}
}
