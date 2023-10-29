package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchFiles(targets, patterns []string, flags options, writer *bufio.Writer) {
	for _, file := range targets {
		findPatterns(file, patterns, flags, writer)
	}
}

func findPatterns(target string, patterns []string, flags options,
	writer *bufio.Writer) {
	bytes, err := os.ReadFile(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: %s: No such file or directory\n", target)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	if flags.c {
		countLines(target, lines, patterns, strings.Contains, flags, writer)
	} else if flags.F {
		applyFlags(target, lines, patterns, func(s1, s2 string) bool { return s1 == s2 }, flags, writer)
	} else {
		applyFlags(target, lines, patterns, strings.Contains, flags, writer)
	}
}

func applyFlags(target string, lines, patterns []string,
	match compare, flags options, writer *bufio.Writer) {
	i := 0
	first := true
lines:
	for i < len(lines) {
		j := 0
		line := lines[i]
		if flags.i {
			line = strings.ToLower(line)
		}
		for j < len(patterns) {
			if match(line, patterns[j]) {
				if !flags.v {
					printLines(target, lines, &i, &first, flags, writer)
				} else {
					i++
				}
				continue lines
			}
			j++
		}
		if flags.v {
			printLines(target, lines, &i, &first, flags, writer)
			continue lines
		}
		i++
	}
}

func countLines(target string, lines, patterns []string,
	match compare, flags options, writer *bufio.Writer) {
	i := 0
	count := 0
lines:
	for i < len(lines) {
		j := 0
		want := !flags.v
		line := lines[i]
		if flags.i {
			line = strings.ToLower(line)
		}
		for j < len(patterns) {
			if match(line, patterns[j]) == want && !flags.v {
				count++
				i++
				continue lines
			}
			j++
		}
		if flags.v {
			count++
		}
		i++
	}
	fmt.Fprintf(writer, "%s:\n%d\n", target, count)
}
