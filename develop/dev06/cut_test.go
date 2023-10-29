package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestOneColumn(t *testing.T) {
	file := "test.txt"
	flags := options{}
	argv := []string{"myCut", "-f", "7", file, "-d", " "}
	files := parseArgs(argv, &flags)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	cutFiles(files, writer, flags)
	writer.Flush()

	want := "Jack\n"
	if buffer.String() != want {
		t.Fatalf(`Wrong output - want: "%s", have: "%s"`, want, buffer.String())
	}
}

func TestMultipleColumns(t *testing.T) {
	file := "test.txt"
	flags := options{}
	argv := []string{"myCut", "-f", "-2", "-f", "4,5", "-f", "7-8", "-f", "9-", file, "-d", " "}
	files := parseArgs(argv, &flags)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	cutFiles(files, writer, flags)
	writer.Flush()

	want := "All work no play Jack a dull boy\n"
	if buffer.String() != want {
		t.Fatalf(`Wrong output - want: "%s", have: "%s"`, want, buffer.String())
	}
}

func TestOtherSeparator(t *testing.T) {
	file := "test.csv"
	flags := options{}
	argv := []string{"myCut", "-f", "-2", "-f", "4,5", "-f", "7-8", "-f", "9-", file, "-d", ";"}
	files := parseArgs(argv, &flags)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	cutFiles(files, writer, flags)
	writer.Flush()

	want := "All work no play Jack a dull boy\n"
	if buffer.String() != want {
		t.Fatalf(`Wrong output - want: "%s", have: "%s"`, want, buffer.String())
	}
}
