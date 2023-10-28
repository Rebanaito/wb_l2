package main

import "testing"

func TestRegularSort(t *testing.T) {
	lines := []string{"cab", "abc", "bac"}
	flags := options{}
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "abc" || lines[1] != "bac" || lines[2] != "cab" {
		t.Fatal("Failed sort (regular)")
	}
}

func TestReverseSort(t *testing.T) {
	lines := []string{"cab", "abc", "bac"}
	flags := options{}
	flags.r = true
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[2] != "abc" || lines[1] != "bac" || lines[0] != "cab" {
		t.Fatal("Failed sort (option -r)")
	}
}

func TestNumSort(t *testing.T) {
	lines := []string{"123", "24", "3"}
	flags := options{}
	flags.n = true
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "3" || lines[1] != "24" || lines[2] != "123" {
		t.Fatal("Failed sort (option -n)")
	}
}

func TestColumnSort(t *testing.T) {
	lines := []string{"1 2 3", "2 3 1", "3 1 2"}
	flags := options{}
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "1 2 3" || lines[1] != "2 3 1" || lines[2] != "3 1 2" {
		t.Fatal("Failed sort (regular)")
	}

	flags.k = 2
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "3 1 2" || lines[1] != "1 2 3" || lines[2] != "2 3 1" {
		t.Fatal("Failed sort (option -k)")
	}

	flags.k = 3
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "2 3 1" || lines[1] != "3 1 2" || lines[2] != "1 2 3" {
		t.Fatal("Failed sort (option -k)")
	}
}

func TestMonthSort(t *testing.T) {
	lines := []string{"Aug 12", "Jan 3", "Oct 26"}
	flags := options{}
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "Aug 12" || lines[1] != "Jan 3" || lines[2] != "Oct 26" {
		t.Fatal("Failed sort (regular)")
	}

	flags.M = true
	quicksort(lines, 0, len(lines)-1, flags)
	if lines[0] != "Jan 3" || lines[1] != "Aug 12" || lines[2] != "Oct 26" {
		t.Fatal("Failed sort (option -k)")
	}
}

func TestUniqueFilter(t *testing.T) {
	lines := []string{"same", "unique", "same", "same"}
	flags := options{}
	quicksort(lines, 0, len(lines)-1, flags)
	lines = removeDuplicates(lines)
	if len(lines) != 2 {
		t.Fatal("Repeating lines not being filtered")
	}
}
