package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(argv []string, flags *options) (lines []string, filename string) {
	filename, err := parseArgs(argv, flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines = strings.Split(string(bytes), "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return
}

func parseArgs(argv []string, flags *options) (filename string, err error) {
	i := 1
	for i < len(argv) {
		switch argv[i] {
		case "-k":
			if i == len(argv)-1 {
				return "", errors.New("-k requires a numeric value > 0")
			}
			num, err := strconv.Atoi(argv[i+1])
			if err != nil || num <= 0 {
				return "", errors.New("-k requires a numeric value > 0")
			}
			flags.k = uint(num)
			i += 1
		case "-o":
			if i == len(argv)-1 {
				return "", errors.New("-o needs a string argument")
			}
			flags.o = argv[i+1]
			i += 1
		case "-n":
			flags.n = true
		case "-r":
			flags.r = true
		case "-u":
			flags.u = true
		case "-M":
			flags.M = true
		case "-c":
			flags.c = true
		default:
			if argv[i][0] == '-' {
				return "", errors.New("unrecognized option: " + argv[i])
			}
			if filename != "" {
				return "", errors.New("incorrect format")
			}
			filename = argv[i]
		}
		i += 1
	}
	if filename == "" {
		err = errors.New("incorrect format")
	}
	return
}
