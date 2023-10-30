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

	"github.com/fatih/color"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	clearScreen()
	for {
		printCommandPrompt()
		input, err := reader.ReadString('\n')
		if exitPrompted(input, err) {
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
		err := command.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func processPipes(input string) {
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

func printCommandPrompt() {
	username, _ := user.Current()
	hostname, _ := os.Hostname()
	pwd, _ := os.Getwd()
	fmt.Printf("%s%s%s%s%s%s", color.HiCyanString("┌──("), color.HiYellowString(username.Username+"@"+hostname),
		color.HiCyanString(")-["), color.HiBlueString(pwd), color.HiCyanString("]\n└─"), color.HiYellowString("$ "))
}

func clearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func exitPrompted(input string, err error) bool {
	return err == io.EOF || input == "\\exit\n"
}
