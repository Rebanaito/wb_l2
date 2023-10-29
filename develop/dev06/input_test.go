package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestValidInput1(t *testing.T) {
	file := "test.txt"
	flags := options{}
	wantOptions := options{d: " ", fieldsList: []int{7}}
	argv := []string{"myCut", "-f", "7", "-d", " ", file, "-s"}
	files := parseArgs(argv, &flags)
	if len(files) != 1 {
		t.Fatal(`Expecting 1 file, have`, len(files))
	}
	if flags.d != " " || !flags.s || len(flags.fieldsList) != 1 || flags.fieldsList[0] != 7 {
		t.Fatal(`Error parsing options, want:`, wantOptions, `got:`, flags)
	}
}

func optionsMatch(have, want options) bool {
	if have.d != want.d || have.s != want.s {
		return false
	}
	if have.fieldsBottom != want.fieldsBottom || have.fieldsTop != want.fieldsTop {
		return false
	}
	if len(have.fieldsList) != len(want.fieldsList) {
		return false
	}
	for i, field := range have.fieldsList {
		if field != want.fieldsList[i] {
			return false
		}
	}
	return true
}

func TestValidInput2(t *testing.T) {
	file := "test.txt"
	flags := options{}
	wantOptions := options{fieldsBottom: 2, fieldsTop: 9, fieldsList: []int{4, 5, 7, 8}, d: "\t"}
	argv := []string{"myCut", "-f", "-2", "-f", "4,5", "-f", "7-8", "-f", "9-", file}
	files := parseArgs(argv, &flags)
	if len(files) != 1 {
		t.Fatal(`Expecting 1 file, have`, len(files))
	}
	if !optionsMatch(flags, wantOptions) {
		t.Fatal(`Error parsing options, want:`, wantOptions, `got:`, flags)
	}
}

func TestInvalidInput1(t *testing.T) {
	if os.Getenv("INVALID_FLAG_CRASH") == "1" {
		file := "test.txt"
		flags := options{}
		argv := []string{"myCut", "-a", file}
		parseArgs(argv, &flags)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidInput1")
	cmd.Env = append(os.Environ(), "INVALID_FLAG_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInvalidInput2(t *testing.T) {
	if os.Getenv("INVALID_FLAG_CRASH") == "1" {
		file := "test.txt"
		flags := options{}
		argv := []string{"myCut", file, "-f"}
		parseArgs(argv, &flags)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidInput2")
	cmd.Env = append(os.Environ(), "INVALID_FLAG_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInvalidInput3(t *testing.T) {
	if os.Getenv("INVALID_FLAG_CRASH") == "1" {
		file := "test.txt"
		flags := options{}
		argv := []string{"myCut", file, "-f", "-3", "-f", "-4"}
		parseArgs(argv, &flags)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidInput3")
	cmd.Env = append(os.Environ(), "INVALID_FLAG_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
