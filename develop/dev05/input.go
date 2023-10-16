package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(argv []string) ([]string, []string, error) {
	var targetFiles []string
	var patternFiles []string
	var patterns []string
	i := 1
	for i < len(argv) {
		switch argv[i] {
		case "-A":
			if i == len(argv)-1 {
				return nil, nil, errors.New("-A requires a non-negative numeric value")
			}
			num, err := strconv.Atoi(argv[i+1])
			if err != nil || num <= 0 {
				return nil, nil, errors.New("-A requires a non-negative numeric value")
			}
			flags.A = uint(num)
			i += 1
		case "-B":
			if i == len(argv)-1 {
				return nil, nil, errors.New("-B requires a non-negative numeric value")
			}
			num, err := strconv.Atoi(argv[i+1])
			if err != nil || num <= 0 {
				return nil, nil, errors.New("-B requires a non-negative numeric value")
			}
			flags.B = uint(num)
			i += 1
		case "-C":
			if i == len(argv)-1 {
				return nil, nil, errors.New("-C requires a non-negative numeric value")
			}
			num, err := strconv.Atoi(argv[i+1])
			if err != nil || num <= 0 {
				return nil, nil, errors.New("-C requires a non-negative numeric value")
			}
			flags.C = uint(num)
			i += 1
		case "-e":
			if i == len(argv)-1 {
				return nil, nil, errors.New("-e requires a string argument")
			}
			patterns = append(patterns, argv[i+1])
			i += 1
		case "-f":
			if i == len(argv)-1 {
				return nil, nil, errors.New("-f requires a string argument")
			}
			patterns = append(patternFiles, argv[i+1])
			i += 1
		case "-c":
			flags.c = true
		case "-i":
			flags.i = true
		case "-v":
			flags.v = true
		case "-F":
			flags.F = true
		case "-n":
			flags.n = true
		default:
			if argv[i][0] == '-' {
				return nil, nil, errors.New("unrecognized option: " + argv[i])
			}
			targetFiles = append(targetFiles, argv[i])
		}
		i += 1
	}
	if len(targetFiles) != 0 {
		parseFiles(&patterns, patternFiles)
	}
	return targetFiles, patterns, nil
}

func parseFiles(patterns *[]string, files []string) {
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "grep: %s: No such file or directory\n", file)
			continue
		}
		lines := strings.Split(string(bytes), "\n")
		lines = lines[:len(lines)-2]
		*patterns = append(*patterns, lines...)
	}
}
