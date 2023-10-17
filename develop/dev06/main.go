package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type options struct {
	fieldsBottom int
	fieldsTop    int
	fieldsList   []int
	d            string
	s            bool
}

var flags options

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	files := parseArgs(os.Args)
	if invalidValues() {
		fmt.Fprintf(os.Stderr, "Invalid field values")
		os.Exit(1)
	}
	if len(files) != 0 {
		cutFiles(files, writer)
	} else {
		cutStdin(writer)
	}
}

func invalidValues() bool {
	if flags.fieldsTop != 0 && flags.fieldsTop < flags.fieldsBottom {
		return true
	}
	if flags.fieldsList != nil {
		var unique []int
		unique = append(unique, flags.fieldsList[0])
		for i := 1; i < len(flags.fieldsList); i++ {
			if flags.fieldsList[i] != flags.fieldsList[i-1] {
				unique = append(unique, flags.fieldsList[i])
			}
		}
		sort.Ints(unique)
		flags.fieldsList = unique
	}
	return false
}

func cutStdin(writer *bufio.Writer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		cutLine(line, writer)
	}
}

func cutLine(line string, writer *bufio.Writer) {
	if !strings.Contains(line, flags.d) {
		if flags.s {
			return
		}
		fmt.Fprint(writer, line)
	} else {
		fields := strings.Split(line, flags.d)
	}
}
