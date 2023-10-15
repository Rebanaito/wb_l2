package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type flags struct {
	k uint
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

func main() {
	flags := flags{}
	argv := os.Args
	filename, err := parseArgs(argv, &flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(flags)
	fmt.Println(filename)
}

func parseArgs(argv []string, flags *flags) (string, error) {
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
