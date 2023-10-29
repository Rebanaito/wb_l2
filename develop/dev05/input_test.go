package main

import (
	"testing"
)

func TestValidInput1(t *testing.T) {
	pattern := "pattern"
	file := "file.txt"
	argv := []string{"myGrep", pattern, file}
	flags := options{}
	blankOptions := options{}
	targets, patterns, err := parseArgs(argv, &flags)
	if err != nil {
		t.Fatal(`Error where not expected`)
	}
	if flags != blankOptions {
		t.Fatal(`Found options where none were present`)
	}
	if len(targets) != 1 || targets[0] != file {
		t.Fatal(`Error parsing target`)
	}
	if len(patterns) != 1 || patterns[0] != pattern {
		t.Fatal(`Error parsing pattern`)
	}
}

func TestValidInput2(t *testing.T) {
	pattern := "pattern"
	files := []string{"file2.txt", "file2.txt"}
	argv := []string{"myGrep", pattern, files[0], files[1]}
	flags := options{}
	blankOptions := options{}
	targets, patterns, err := parseArgs(argv, &flags)
	if err != nil {
		t.Fatal(`Error where not expected`)
	}
	if flags != blankOptions {
		t.Fatal(`Found options where none were present`)
	}
	if len(targets) != 2 || targets[0] != files[0] || targets[1] != files[1] {
		t.Fatal(`Error parsing targets`)
	}
	if len(patterns) != 1 || patterns[0] != pattern {
		t.Fatal(`Error parsing pattern`)
	}
}

func TestValidInput3(t *testing.T) {
	pattern := "pattern"
	file := "file.txt"
	argv := []string{"myGrep", pattern, "-n", "-B", "3", "-c", "-A", "2", "-i", file, "-v", "-F", "-C", "4"}
	flags := options{}
	wantOptions := options{n: true, B: 3, c: true, A: 2, i: true, v: true, F: true, C: 4}
	targets, patterns, err := parseArgs(argv, &flags)
	if err != nil {
		t.Fatal(`Error where not expected`)
	}
	if flags != wantOptions {
		t.Fatalf(`Error parsing options - want: %v, got: %v`, wantOptions, flags)
	}
	if len(targets) != 1 || targets[0] != file {
		t.Fatal(`Error parsing target`)
	}
	if len(patterns) != 1 || patterns[0] != pattern {
		t.Fatal(`Error parsing pattern`)
	}
}

func TestPatternFlags(t *testing.T) {
	patt := []string{"pattern1", "pattern2"}
	files := []string{"patternFile1.txt", "patternFile2.txt"}
	argv := []string{"myGrep", "-f", files[0], "-e", patt[0], "file.txt", "-f", files[1], "-e", patt[1]}
	flags := options{}
	blankOptions := options{}
	targets, patterns, err := parseArgs(argv, &flags)
	if err != nil {
		t.Fatal(`Error where not expected`)
	}
	if flags != blankOptions {
		t.Fatal(`Found options where none were present`)
	}
	if len(targets) != 1 || targets[0] != "file.txt" {
		t.Fatal(`Error parsing targets (option -f)`)
	}
	if len(patterns) != 2 || patterns[0] != patt[0] || patterns[1] != patt[1] {
		t.Fatal(`Error parsing patterns (option -e)`)
	}
}

func TestInvalidInput(t *testing.T) {
	flags := options{}

	argv := []string{"myGrep", "-s", "pattern", "test.txt"}
	want := "unrecognized option: -s"
	_, _, err := parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "pattern"}
	want = "invalid format"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "-e", "pattern", "-e", "pattern2"}
	want = "invalid format"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "file.txt", "-e"}
	want = "-e requires a string argument"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "file.txt", "-f"}
	want = "-f requires a string argument"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "-f", "patternFile.txt", "-e", "pattern"}
	want = "invalid format"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "pattern", "file.txt", "-A"}
	want = "-A requires a non-negative numeric value"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "pattern", "file.txt", "-B"}
	want = "-B requires a non-negative numeric value"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}

	argv = []string{"myGrep", "pattern", "file.txt", "-C"}
	want = "-C requires a non-negative numeric value"
	_, _, err = parseArgs(argv, &flags)
	if err == nil || err.Error() != want {
		t.Fatalf(`Expecting error: "%s", got: "%s"`, want, err.Error())
	}
}
