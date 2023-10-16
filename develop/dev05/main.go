package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type options struct {
	A uint
	B uint
	C uint
	c bool
	i bool
	v bool
	F bool
	n bool
}

var flags options

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	targets, patterns, err := parseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, target := range targets {
		findPatterns(target, patterns, writer)
	}
}

func findPatterns(target string, patterns []string, writer *bufio.Writer) {
	first := true
	bytes, err := os.ReadFile(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: %s: No such file or directory\n", target)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-2]
	i := 0
	for i < len(lines) {
		j := 0
		contains := false
		for !contains && j < len(patterns) {
			if strings.Contains(lines[i], patterns[j]) {
				printLines(target, lines, &i, &first, writer)
			}
			j++
		}
	}
}

func printLines(target string, lines []string, index *int, first *bool, writer *bufio.Writer) {
	if *first {
		fmt.Fprintf(writer, "%s:\n", target)
		*first = false
	}
	start := *index
	end := *index

}
