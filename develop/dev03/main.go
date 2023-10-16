package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type options struct {
	k uint
	n bool
	r bool
	u bool
	M bool
	c bool
	o string
}

var flags options

var month = map[string]int{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func main() {
	argv := os.Args
	filename, err := parseArgs(argv)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-2]
	if flags.c {
		sorted := checkIfSorted(lines)
		if sorted {
			fmt.Println("Lines in file '" + filename + "' are sorted")
		} else {
			fmt.Println("Lines in file '" + filename + "' are not sorted")
			os.Exit(1)
		}
	} else {
		quicksort(lines, 0, len(lines)-1)
		outfile := "output.txt"
		if flags.o != "" {
			outfile = flags.o
		}
		file, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening the output file")
			os.Exit(1)
		}
		for _, line := range lines {
			fmt.Fprintln(file, line)
		}
	}
}

func quicksort(lines []string, low, high int) {
	if low >= high {
		return
	}
	pivot := quicksortPartition(lines, low, high)
	quicksort(lines, low, pivot-1)
	quicksort(lines, pivot+1, high)
}

func quicksortPartition(lines []string, low, high int) int {
	pivot := lines[high]
	i := low - 1
	for j := low; j < high; j++ {
		if compare(lines[j], pivot) {
			i++
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	lines[i+1], lines[high] = lines[high], lines[i+1]
	return i + 1
}

func checkIfSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if !compare(lines[i-1], lines[i]) {
			return false
		}
	}
	return true
}

func compare(left, right string) bool {
	var result bool
	if flags.k != 0 {
		result = substring(left, right)
	} else if flags.n {
		result = numbers(left, right)
	} else if flags.M {
		result = months(left, right)
	} else {
		result = left < right
	}
	if flags.r {
		result = !result
	}
	return result
}

func substring(l, r string) bool {
	left := strings.Split(l, " ")
	right := strings.Split(r, " ")
	if len(left) < int(flags.k) && len(right) >= int(flags.k) {
		return true
	} else if len(right) < int(flags.k) && len(left) >= int(flags.k) {
		return false
	} else if len(left) < int(flags.k) && len(right) < int(flags.k) {
		return left[0] < right[0]
	}
	return left[flags.k-1] < right[flags.k-1]
}

func numbers(l, r string) bool {
	left := parseInt(l)
	right := parseInt(r)
	return left < right
}

func parseInt(str string) int {
	chars := []rune(str)
	i := 0
	for i < len(chars) {
		if unicode.IsSpace(chars[i]) {
			i++
			continue
		} else {
			break
		}
	}
	negative := false
	if i < len(chars) && chars[i] == '-' {
		negative = true
		i++
	} else if i < len(chars) && chars[i] == '+' {
		i++
	}
	num := 0
	for i < len(chars) {
		if chars[i] >= '0' && chars[i] <= '9' {
			num = num*10 + int(chars[i]-'0')
			i++
		} else {
			break
		}
	}
	if negative {
		num = -num
	}
	return num
}

func months(l, r string) bool {
	left := firstThree(l)
	right := firstThree(r)
	if left == 0 && right == 0 {
		return l < r
	}
	return left < right
}

func firstThree(str string) int {
	var builder strings.Builder
	count := 3
	for _, char := range str {
		if unicode.IsSpace(char) {
			continue
		}
		if char >= 'a' && char <= 'z' {
			builder.WriteRune(char - 32)
		} else {
			builder.WriteRune(char)
		}
		count--
		if count == 0 {
			break
		}
	}
	return month[builder.String()]
}
