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

var flags options

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
	lines, filename := parseInput(os.Args)
	if flags.c {
		check(lines, filename)
	} else {
		sortLines(lines)
	}
}

func sortLines(lines []string) {
	quicksort(lines, 0, len(lines)-1)
	outfile := "output.txt"
	if flags.o != "" {
		outfile = flags.o
	}
	file, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the output file")
		os.Exit(1)
	}
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
}
