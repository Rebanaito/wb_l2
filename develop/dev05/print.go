package main

import (
	"bufio"
	"fmt"
)

func printLines(target string, lines []string, index *int,
	first *bool, flags options, writer *bufio.Writer) {
	if *first {
		fmt.Fprintf(writer, "%s:\n", target)
		*first = false
	}
	start := *index
	end := start
	rangesToPrint(&start, &end, len(lines), flags)
	writeToBuffer(lines, start, end, flags, writer)
	*index = end + 1
}

func rangesToPrint(start, end *int, size int, flags options) {
	if flags.C != 0 {
		*start -= int(flags.C)
		*end += int(flags.C)
	}
	if flags.B > flags.C {
		*start -= int(flags.B - flags.C)
	}
	if flags.A > flags.C {
		*end += int(flags.A - flags.C)
	}
	if *start < 0 {
		*start = 0
	}
	if *end >= size {
		*end = size - 1
	}
}

func writeToBuffer(lines []string, start, end int, flags options, writer *bufio.Writer) {
	for start <= end {
		if flags.n {
			fmt.Fprintf(writer, "%d: ", start+1)
		}
		fmt.Fprintln(writer, lines[start])
		start++
	}
}
