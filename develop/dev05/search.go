package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findPatterns(target string, patterns []string, flags options, writer *bufio.Writer) {
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
		applyFlags(target, lines, patterns, &first, func(s1, s2 string) bool { return s1 == s2 }, flags, writer)
	} else {
		applyFlags(target, lines, patterns, &first, strings.Contains, flags, writer)
	}
}

func applyFlags(target string, lines, patterns []string, first *bool, match compare, flags options, writer *bufio.Writer) {
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
				printLines(target, lines, &i, first, flags, writer)
				continue lines
			}
			j++
		}
		i++
	}
}
