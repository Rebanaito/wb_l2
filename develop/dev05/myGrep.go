package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	writer := bufio.NewWriter(os.Stdout)
	flags := options{}
	defer writer.Flush()
	targets, patterns, err := parseArgs(os.Args, &flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	searchFiles(targets, patterns, flags, writer)
}
