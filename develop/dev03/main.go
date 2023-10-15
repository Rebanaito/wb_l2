package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type options struct {
	k uint
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

var flags options

func main() {
	argv := os.Args
	filename, err := parseArgs(argv, &flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")
	if flags.c {
		sorted := checkIfSorted(lines)
		if sorted {
			fmt.Println("Lines in file '" + filename + "' are sorted")
		} else {
			fmt.Println("Lines in file '" + filename + "' are not sorted")
			os.Exit(1)
		}
	} else {
		quicksort(lines, 0, len(lines)-1)
	}
}

func parseArgs(argv []string, flags *options) (string, error) {
	var filename string
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
			i += 2
		case "-n":
			flags.n = true
			i += 1
		case "-r":
			flags.r = true
			i += 1
		case "-u":
			flags.u = true
			i += 1
		case "-M":
			flags.M = true
			i += 1
		case "-b":
			flags.b = true
			i += 1
		case "-c":
			flags.c = true
			i += 1
		case "-h":
			flags.h = true
			i += 1
		default:
			if argv[i][0] == '-' {
				return "", errors.New("unrecognized option: " + argv[i])
			}
			if filename != "" {
				return "", errors.New("incorrect format")
			}
			filename = argv[i]
			i += 1
		}
	}
	return filename, nil
}

func quicksort(lines []string, low, high int) {
	if low >= high {
		return
	}
	pivot := quicksortPartition(lines, low, high)
	quicksort(lines, low, pivot-1)
	quicksort(lines, pivot+1, high)
}

func quicksortPartition(lines []string, low, high int) int {
	pivot := lines[high]
	i := low - 1
	for j := low; j < high; j++ {
		if compare(lines[j], pivot) {
			i++
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	lines[i+1], lines[high] = lines[high], lines[i+1]
	return i + 1
}

func checkIfSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if !compare(lines[i-1], lines[i]) {
			return false
		}
	}
	return true
}

func compare(left, right string) bool {
	var result bool
	if flags.k != 0 {

	} else if flags.n {

	} else if flags.M {

	} else if flags.h {

	} else {

	}
	if flags.r {
		result = !result
	}
	return result
}
