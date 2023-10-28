package main

import (
	"fmt"
	"log"
)

func check(lines []string, filename string) {
	sorted := checkIfSorted(lines)
	if sorted {
		fmt.Println("Lines in file '" + filename + "' are sorted")
	} else {
		log.Fatal("Lines in file '" + filename + "' are not sorted")
	}
}

func checkIfSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if !compare(lines[i-1], lines[i]) {
			return false
		}
	}
	return true
}
