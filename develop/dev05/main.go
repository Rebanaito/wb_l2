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

type compare func(string, string) bool

var flags options
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	targets, patterns, err := parseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, target := range targets {
		findPatterns(target, patterns)
	}
}

func findPatterns(target string, patterns []string) {
	first := true
	bytes, err := os.ReadFile(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: %s: No such file or directory\n", target)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	if flags.F {
		applyFlags(target, lines, patterns, &first, func(s1, s2 string) bool { return s1 == s2 })
	} else {
		applyFlags(target, lines, patterns, &first, strings.Contains)
	}
}

func applyFlags(target string, lines, patterns []string, first *bool, match compare) {
	i := 0
lines:
	for i < len(lines) {
		j := 0
		want := !flags.v
		line := lines[i]
		if flags.i {
			line = strings.ToLower(line)
		}
		for j < len(patterns) {
			if match(line, patterns[j]) == want {
				printLines(target, lines, &i, first)
				continue lines
			}
			j++
		}
		i++
	}
}

func printLines(target string, lines []string, index *int, first *bool) {
	if *first {
		fmt.Fprintf(writer, "%s:\n", target)
		*first = false
	}
	start := *index
	end := *index
	if flags.C != 0 {
		start -= int(flags.C)
		end += int(flags.C)
	}
	if flags.B != 0 {
		start -= int(flags.B)
	}
	if flags.A != 0 {
		end += int(flags.A)
	}
	if start < 0 {
		start = 0
	}
	if end >= len(lines) {
		end = len(lines) - 1
	}
	for start <= end {
		if flags.n {
			fmt.Fprintf(writer, "%d: ", start+1)
		}
		fmt.Fprintln(writer, lines[start])
		start++
	}
	*index = end + 1
}
