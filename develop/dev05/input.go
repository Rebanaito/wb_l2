package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(argv []string, flags *options) (targetFiles, patterns []string, err error) {
	var patternFiles []string
	i := 1
	for i < len(argv) {
		switch argv[i] {
		case "-A", "-B", "-C":
			err := flagsRange(argv[i], argv, &i, flags)
			if err != nil {
				return nil, nil, err
			}
		case "-e", "-f":
			err := flagsPatterns(argv[i], argv, &patterns, &patternFiles, &i, flags)
			if err != nil {
				return nil, nil, err
			}
		case "-c", "-i", "-v", "-F", "-n":
			flagsBool(argv[i], flags)
		default:
			if argv[i][0] == '-' {
				return nil, nil, errors.New("unrecognized option: " + argv[i])
			}
			targetFiles = append(targetFiles, argv[i])
		}
		i += 1
	}
	err = collectPatterns(&patterns, &targetFiles, patternFiles)
	if err != nil {
		return nil, nil, err
	}
	caseInsensitive(patterns, flags.i)
	return
}

func flagsRange(flag string, argv []string, i *int, flags *options) error {
	if *i == len(argv)-1 {
		return errors.New(flag + " requires a non-negative numeric value")
	}
	num, err := strconv.Atoi(argv[*i+1])
	if err != nil || num <= 0 {
		return errors.New(flag + " requires a non-negative numeric value")
	}
	var field *uint
	switch flag {
	case "-A":
		field = &flags.A
	case "-B":
		field = &flags.B
	case "-C":
		field = &flags.C
	}
	*field = uint(num)
	*i += 1
	return nil
}

func flagsPatterns(flag string, argv []string, patterns,
	patternFiles *[]string, i *int, flags *options) error {
	if *i == len(argv)-1 {
		return errors.New(flag + " requires a string argument")
	}
	switch flag {
	case "-e":
		*patterns = append(*patterns, argv[*i+1])
	case "-f":
		*patternFiles = append(*patternFiles, argv[*i+1])
	}
	*i += 1
	return nil
}

func flagsBool(flag string, flags *options) {
	switch flag {
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
	}
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

func caseInsensitive(patterns []string, insensitive bool) {
	if insensitive {
		for i := range patterns {
			patterns[i] = strings.ToLower(patterns[i])
		}
	}
}

func collectPatterns(patterns, targetFiles *[]string, patternFiles []string) error {
	if len(patternFiles) != 0 {
		parseFiles(patterns, patternFiles)
	}
	if len(*patterns) == 0 && len(*targetFiles) > 1 {
		*patterns = append(*patterns, (*targetFiles)[0])
		*targetFiles = (*targetFiles)[1:]
	} else if len(*patterns) == 0 || len(*targetFiles) == 0 {
		return errors.New("invalid format")
	}
	return nil
}
