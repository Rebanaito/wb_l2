package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(argv []string) []string {
	var files []string
	flags.d = "\t"
	for i := 1; i < len(argv); i++ {
		switch argv[i] {
		case "-f":
			if i == len(argv)-1 {
				fmt.Fprint(os.Stderr, "-f: needs field number(s) as argument(s)\n")
				os.Exit(1)
			}
			parseList(argv[i+1])
		case "-d":
			if i == len(argv)-1 {
				fmt.Fprint(os.Stderr, "-d: needs a string as an argument)\n")
				os.Exit(1)
			}
			flags.d = argv[i+1]
		case "-s":
			flags.s = true
		default:
			if argv[i][0] == '-' {
				fmt.Fprintf(os.Stderr, "Unrecognizer option: %s\n", argv[i])
				os.Exit(1)
			}
			files = append(files, argv[i])
		}
	}
	return files
}

func parseList(list string) {
	if strings.ContainsRune(list, '-') {
		parseRange(list)
	} else if strings.ContainsRune(list, ',') {
		parseFields(list)
	} else {
		num, err := strconv.Atoi(list)
		if err != nil || num < 1 {
			fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", list)
		}
		flags.fieldsList = append(flags.fieldsList, num)
	}
}

func parseRange(str string) {
	if str[0] == '-' {
		parseBottom(str)
	} else if str[len(str)-1] == '-' {
		parseTop(str)
	} else {
		halves := strings.Split(str, "-")
		if len(halves) != 2 {
			fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", str)
			os.Exit(1)
		}
		left, err1 := strconv.Atoi(halves[0])
		right, err2 := strconv.Atoi(halves[1])
		if err1 != nil || err2 != nil || left > right {
			fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", str)
			os.Exit(1)
		}
		rnge := make([]int, right-left+1)
		for i := 0; i < len(rnge); i++ {
			rnge[i] = left + i
		}
		flags.fieldsList = append(flags.fieldsList, rnge...)
	}
}

func parseBottom(str string) {
	if flags.fieldsBottom != 0 {
		fmt.Fprint(os.Stderr, "-f: too many range specifiers\n")
		os.Exit(1)
	}
	bottom, err := strconv.Atoi(str[:len(str)-1])
	if err != nil || bottom < 1 {
		fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", str)
		os.Exit(1)
	}
	flags.fieldsBottom = bottom
}

func parseTop(str string) {
	if flags.fieldsTop != 0 {
		fmt.Fprint(os.Stderr, "-f: too many range specifiers\n")
		os.Exit(1)
	}
	top, err := strconv.Atoi(str[1:])
	if err != nil || top < 1 {
		fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", str)
		os.Exit(1)
	}
	flags.fieldsTop = top
}

func parseFields(fields string) {
	substrings := strings.Split(fields, ",")
	if len(substrings) == 0 {
		fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", fields)
		os.Exit(1)
	}
	var nums []int
	for _, str := range substrings {
		num, err := strconv.Atoi(str)
		if err != nil || num < 1 {
			fmt.Fprintf(os.Stderr, "-f: invalid argument: %s\n", str)
			os.Exit(1)
		}
		nums = append(nums, num)
	}
	flags.fieldsList = append(flags.fieldsList, nums...)
}
