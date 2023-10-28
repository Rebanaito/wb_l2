package main

import (
	"fmt"
	"os"
)

type options struct {
	k uint
	n bool
	r bool
	u bool
	M bool
	c bool
	o string
}

var month = map[string]int{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func main() {
	var flags options
	lines, filename := parseInput(os.Args, &flags)
	if flags.c {
		check(lines, filename, flags)
	} else {
		quicksort(lines, 0, len(lines)-1, flags)
		writeToFile(lines, flags)
	}
}

func writeToFile(lines []string, flags options) {
	outfile := "output.txt"
	if flags.o != "" {
		outfile = flags.o
	}
	file, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the output file")
		os.Exit(1)
	}
	defer file.Close()
	if flags.u {
		lines = removeDuplicates(lines)
	}
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
}

func removeDuplicates(lines []string) (new []string) {
	if len(lines) != 0 {
		new = append(new, lines[0])
	}
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			new = append(new, lines[i])
		}
	}
	return
}
