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

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	flags := options{}
	files := parseArgs(os.Args, &flags)
	if invalidValues(flags) {
		fmt.Fprintf(os.Stderr, "Invalid field values")
		os.Exit(1)
	}
	if len(files) != 0 {
		cutFiles(files, writer, flags)
	} else {
		cutStdin(writer, flags)
	}
}

func invalidValues(flags options) bool {
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

func cutStdin(writer *bufio.Writer, flags options) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		cutLine(line, writer, flags)
	}
}

func cutFiles(files []string, writer *bufio.Writer, flags options) {
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cut: %s: No such file or directory\n", file)
			continue
		}
		lines := strings.Split(string(bytes), "\n")
		for _, line := range lines {
			cutLine(line, writer, flags)
		}
	}
}

func cutLine(line string, writer *bufio.Writer, flags options) {
	if !strings.Contains(line, flags.d) {
		if flags.s {
			return
		}
		fmt.Fprint(writer, line)
	} else {
		fields := strings.Split(line, flags.d)
		str := buildString(fields, flags)
		fmt.Fprintln(writer, str)
	}
}

func buildString(fields []string, flags options) string {
	var builder strings.Builder
	i := 0
	first := true
	if flags.fieldsBottom != 0 {
		for i < flags.fieldsBottom && i < len(fields) {
			stringAppend(&first, &builder, fields, i, flags)
			i++
		}
	}
	if flags.fieldsList != nil {
		for j := 0; j < len(flags.fieldsList) &&
			(flags.fieldsTop == 0 || flags.fieldsList[j] < flags.fieldsTop); j++ {
			if flags.fieldsList[j]-1 <= i && !first {
				continue
			} else if flags.fieldsList[j]-1 >= len(fields) {
				break
			}
			stringAppend(&first, &builder, fields, flags.fieldsList[j]-1, flags)
		}
	}
	if flags.fieldsTop != 0 {
		k := flags.fieldsTop - 1
		for k < len(fields) {
			stringAppend(&first, &builder, fields, k, flags)
			k++
		}
	}
	return builder.String()
}

func stringAppend(first *bool, builder *strings.Builder, fields []string, index int, flags options) {
	if !*first {
		builder.WriteString(flags.d)
	} else {
		*first = false
	}
	builder.WriteString(fields[index])
}
