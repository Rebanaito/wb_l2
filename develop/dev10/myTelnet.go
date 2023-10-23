package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var TIMEOUT int = 10

func main() {
	host, port, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	conn, err := net.DialTimeout("tcp", host+":"+port, time.Second*time.Duration(TIMEOUT))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	takeInput(conn)
}

func parseArgs(argv []string) (host, port string, err error) {
	for _, arg := range argv {
		if strings.Contains(arg, "--timeout=") {
			err := parseTimeout(arg)
			if err != nil {
				return "", "", err
			}
		} else {
			if host == "" {
				host = arg
			} else if port == "" {
				port = arg
			} else {
				return "", "", fmt.Errorf("telnet: too many arguments")
			}
		}
	}
	return host, port, nil
}

func parseTimeout(arg string) (err error) {
	arg = arg[10:]
	lastChar := arg[len(arg)-1]
	var num int
	if !(lastChar >= '0' && lastChar <= '9') {
		num, err = strconv.Atoi(arg[:len(arg)-1])
		switch lastChar {
		case 'm':
			num *= 60
		case 'h':
			num *= 3600
		case 's':
			break
		default:
			return fmt.Errorf("telnet: --timeout: invalid argument")
		}
	} else {
		num, err = strconv.Atoi(arg)
	}
	if err != nil || num < 0 {
		return err
	}
	TIMEOUT = num
	return err
}

func takeInput(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	reply := make([]byte, 8096)
	for {
		printPrompt()
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		input = input[:len(input)-1]
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		_, err = conn.Read(reply)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, err)
			}
			break
		}
		fmt.Fprintln(writer, string(reply))
		writer.Flush()
	}
}

func printPrompt() {
	fmt.Print("> ")
}
