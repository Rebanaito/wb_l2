package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF || input == "\\exit\n" {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		processInput(input[:len(input)-1])
		writer.Flush()
	}
}

func processInput(input string) {
	if strings.Contains(input, "|") {

	} else {
		argv := splitWithQuotes(input)
		if len(argv) == 0 {
			return
		}
		processCommand(argv)
	}
}

func processCommand(argv []string) {
	switch argv[0] {
	case "echo":
		echo(argv)
	case "pwd":
		pwd(argv)
	case "cd":
		cd(argv)
	case "kill":
		kill(argv)
	case "ps":
		ps(argv)
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found", argv[0])
	}
}

func pwd(argv []string) {
	if len(argv) != 1 {
		fmt.Fprintf(os.Stderr, "pwd: too many arguments\n")
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(writer, dir)
	}
}

func echo(argv []string) {
	for i := 1; i < len(argv); i++ {
		if i != 1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, argv[i])
	}
	fmt.Fprintln(writer)
}
