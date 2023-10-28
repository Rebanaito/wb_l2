package main

import (
	"testing"
)

func TestNoOption(t *testing.T) {
	flags := options{}
	lines := []string{"abc", "bac", "cab"}
	sorted := checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines`)
	}

}

func TestOptionK(t *testing.T) {
	flags := options{}

	lines := []string{"line 3", "line 1", "line 2"}
	sorted := checkIfSorted(lines, flags)
	if sorted {
		t.Fatal(`Should not pass on unsorted lines`)
	}

	flags.k = 1
	sorted = checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines (option -k)`)
	}
}

func TestOptionN(t *testing.T) {
	flags := options{}
	lines := []string{"3", "12", "24"}
	sorted := checkIfSorted(lines, flags)
	if sorted {
		t.Fatal(`Should not pass on unsorted lines`)
	}

	flags.n = true
	sorted = checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines (option -n)`)
	}
}

func TestOptionM(t *testing.T) {
	flags := options{}
	lines := []string{"Aug 12th", "Jan 3rd", "Jul 24th"}
	sorted := checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines`)
	}

	flags.M = true
	sorted = checkIfSorted(lines, flags)
	if sorted {
		t.Fatal(`Should not pass on unsorted lines (option -M)`)
	}

	lines[0], lines[2] = lines[2], lines[0]
	lines[0], lines[1] = lines[1], lines[0]
	sorted = checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines (option -M)`)
	}
}

func TestOptionR(t *testing.T) {
	flags := options{}
	lines := []string{"cab", "bac", "abc"}
	sorted := checkIfSorted(lines, flags)
	if sorted {
		t.Fatal(`Should not pass on unsorted lines`)
	}

	flags.r = true
	sorted = checkIfSorted(lines, flags)
	if !sorted {
		t.Fatal(`False-flagged sorted lines (option -r)`)
	}
}
