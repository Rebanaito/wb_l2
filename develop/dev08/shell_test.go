package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

func TestValidCommands(t *testing.T) {
	want, _ := os.Getwd()
	input := "pwd"
	buffer := bytes.Buffer{}
	writer = bufio.NewWriter(&buffer)
	processInput(input)
	writer.Flush()
	if buffer.String() != want+"\n" {
		t.Fatalf(`Error executing command "pwd" - want: "%s", got: "%s"`,
			want, buffer.String())
	}

	want = "Hello, world!"
	input = `echo "Hello, world!"`
	buffer = bytes.Buffer{}
	writer = bufio.NewWriter(&buffer)
	processInput(input)
	writer.Flush()
	if buffer.String() != want+"\n" {
		t.Fatalf(`Error executing command "echo" - want: "%s", got: "%s"`,
			want, buffer.String())
	}
}

func TestInvalidCommands(t *testing.T) {
	input := "invalidcommand"
	want := "fork/exec /usr/bin/invalidcommand: no such file or directory"
	buffer := bytes.Buffer{}
	tmp := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	processInput(input)
	w.Close()
	io.Copy(&buffer, r)
	os.Stderr = tmp
	if buffer.String() != want+"\n" {
		t.Fatalf(`Error message missing or not matching - want: "%s", got: "%s"`,
			want, buffer.String())
	}
}
