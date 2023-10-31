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
	"sync"
	"time"
)

var TIMEOUT int = 10
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	host, port, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	dialAndCommunicate(host, port)
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

func dialAndCommunicate(host, port string) {
	conn, err := net.DialTimeout("tcp", host+":"+port, time.Second*time.Duration(TIMEOUT))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go listener(conn, wg)
	go writeToConnection(conn, wg)
	wg.Wait()
}

func listener(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make([]byte, 8096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			fmt.Println("Connection closed on the other end")
			break
		}
		if n != 0 {
			fmt.Fprintln(writer, string(buffer))
			writer.Flush()
		}
	}
}

func writeToConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		printPrompt()
		input, done := takeInput()
		if done {
			break
		}
		_, err := conn.Write([]byte(input[:len(input)-1]))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func takeInput() (input string, done bool) {
	input, err := reader.ReadString('\n')
	if err == io.EOF {
		done = true
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		done = true
	}
	return
}

func printPrompt() {
	fmt.Print("> ")
}
