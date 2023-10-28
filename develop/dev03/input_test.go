package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestInput(t *testing.T) {
	var optionsWant options = options{k: 2, r: true, M: true, o: "out.file"}
	fileWant := "files/test.txt"
	argv := []string{"sort", "-k", "2", "-r", "-M", fileWant, "-o", "out.file"}
	var flags options
	lines, filename := parseInput(argv, &flags)
	if filename != fileWant {
		t.Fatalf(`Filename should be: "%s", got: "%s"`, fileWant, filename)
	}
	if !linesMatch(lines) {
		t.Fatal(`Parsed lines don't match the file`)
	}
	if flags != optionsWant {
		t.Fatalf(`Error in the options parser - want: "%v", got: "%v"`, optionsWant, flags)
	}
}

func linesMatch(lines []string) bool {
	return len(lines) == 3 && lines[0] == "Lorem ipsum" && lines[1] == "sample text" && lines[2] == "loop boop"
}

func TestInvalidFile(t *testing.T) {
	if os.Getenv("NO_FILE_CRASH") == "1" {
		argv := []string{"sort", "-k", "2", "-r", "-M", "invalid.file"}
		var flags options
		parseInput(argv, &flags)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidFile")
	cmd.Env = append(os.Environ(), "NO_FILE_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMissingFile(t *testing.T) {
	if os.Getenv("NO_FILE_CRASH") == "1" {
		argv := []string{"sort", "-k", "2", "-r", "-M"}
		var flags options
		parseInput(argv, &flags)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidFile")
	cmd.Env = append(os.Environ(), "NO_FILE_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestTooManyFiles(t *testing.T) {
	argv := []string{"sort", "file1.txt", "file2.txt"}
	var flags options
	errorWant := "incorrect format"
	_, err := parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}

	argv = []string{"sort", "file.txt", "-a"}
	errorWant = "unrecognized option: -a"
	_, err = parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}
}

func TestInvalidOption(t *testing.T) {
	var flags options
	argv := []string{"sort", "file.txt", "-a"}
	errorWant := "unrecognized option: -a"
	_, err := parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}
}

func TestInvalidK(t *testing.T) {
	var flags options
	argv := []string{"sort", "file.txt", "-k"}
	errorWant := "-k requires a numeric value > 0"
	_, err := parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}

	argv = []string{"sort", "file.txt", "-k", "21a"}
	errorWant = "-k requires a numeric value > 0"
	_, err = parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}
}

func TestInvalidO(t *testing.T) {
	var flags options
	argv := []string{"sort", "file.txt", "-o"}
	errorWant := "-o needs a string argument"
	_, err := parseArgs(argv, &flags)
	if err == nil {
		t.Fatal(`No error where one is expected`)
	}
	if err.Error() != errorWant {
		t.Fatalf(`Want error: "%s", got: "%s"`, errorWant, err.Error())
	}
}
