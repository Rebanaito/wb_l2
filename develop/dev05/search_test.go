package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestSimpleSearch(t *testing.T) {
	patterns := []string{"hello", "world"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	flags := options{}
	file := "test.txt"
	want := "hello, friend\nthe best in the world\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (regular) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestFullMatch(t *testing.T) {
	patterns := []string{"strict policy", "out", "world", "hello"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	flags := options{F: true}
	file := "test.txt"
	want := "strict policy\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (full match) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestRanges(t *testing.T) {
	patterns := []string{"policy"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	flags := options{A: 1}
	file := "test.txt"
	want := "strict policy\nthe best in the world\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -A) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}

	buffer.Reset()
	flags = options{B: 2}
	want = "hello, friend\nbranching out\nstrict policy\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -B) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}

	buffer.Reset()
	flags = options{C: 3}
	want = "hello, friend\nbranching out\nstrict policy\nthe best in the world\nlorem ipsum\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -C) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestCount(t *testing.T) {
	patterns := []string{"friend", "policy", "ipsum"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	file := "test.txt"

	flags := options{c: true}
	want := "3\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -c) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestCaseInsensitive(t *testing.T) {
	pattern := "lOrEm IpSuM"
	file := "test.txt"
	argv := []string{"myGrep", pattern, file}
	flags := options{i: true}
	_, patterns, err := parseArgs(argv, &flags)

	if err != nil {
		t.Fatal(`Error where not expected`)
	}
	want := "lorem ipsum\n"
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -i) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestInverse(t *testing.T) {
	patterns := []string{"friend", "out", "policy", "ipsum"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	file := "test.txt"

	flags := options{v: true}
	want := "the best in the world\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -v) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}

func TestNumbered(t *testing.T) {
	patterns := []string{"friend", "policy", "ipsum"}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	file := "test.txt"

	flags := options{n: true}
	want := "1: hello, friend\n3: strict policy\n5: lorem ipsum\n"
	findPatterns(file, patterns, flags, writer)
	writer.Flush()
	if buffer.String() != file+":\n"+want {
		t.Fatalf(`Wrong output (option -n) - want: "%s", got: "%s"`, file+":\n"+want, buffer.String())
	}
}
