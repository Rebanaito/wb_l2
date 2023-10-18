package main

import (
	"bufio"
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
	clear.Run()
}

func processInput(input string) {
	if strings.Contains(input, "|") {

	} else {
		argv := splitWithQuotes(input)
		if len(argv) == 0 {
			return
		}
		command := exec.Command("/usr/bin/"+argv[0], argv[1:]...)
		command.Stdout = writer
		command.Stderr = os.Stderr
		command.Run()
	}
}
