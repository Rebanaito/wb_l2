package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	for {
		username, _ := user.Current()
		hostname, _ := os.Hostname()
		pwd, _ := os.Getwd()
		fmt.Printf("г---(%s@%s)-[%s]\nL-$ ", username.Username, hostname, pwd)
		input, err := reader.ReadString('\n')
		if err == io.EOF || input == "\\exit\n" {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		processInput(input[:len(input)-1])
		fmt.Fprintln(writer)
		writer.Flush()
	}
}

func processInput(input string) {
	if strings.Contains(input, "|") {
		processPipes(input)
	} else {
		argv := splitWithQuotes(input, ' ')
		if len(argv) == 0 {
			return
		}
		command := exec.Command("/usr/bin/"+argv[0], argv[1:]...)
		command.Stdout = writer
		command.Stderr = os.Stderr
		command.Run()
	}
}

func processPipes(input string) {
	fmt.Println(input)
	pipes := splitWithQuotes(input, '|')
	commands := make([]*exec.Cmd, len(pipes))
	for i := range pipes {
		pipes[i] = strings.Trim(pipes[i], " ")
		argv := splitWithQuotes(pipes[i], ' ')
		if len(argv) == 0 {
			fmt.Fprintln(os.Stderr, "Invalid input")
			return
		}
		commands[i] = exec.Command("/usr/bin/"+argv[0], argv[1:]...)
		commands[i].Stderr = os.Stderr
	}
	buffer, _ := commands[0].Output()
	for i := 1; i < len(commands); i++ {
		commands[i].Stdin = bytes.NewReader(buffer)
		buffer, _ = commands[i].Output()
	}
	fmt.Fprintln(writer, string(buffer))
}

func splitWithQuotes(input string, sep rune) []string {
	reader := csv.NewReader(strings.NewReader(input))
	reader.Comma = sep
	reader.LazyQuotes = true
	args, err := reader.Read()
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil
	}
	return args
}
