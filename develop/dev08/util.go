package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func splitWithQuotes(input string) []string {
	reader := csv.NewReader(strings.NewReader(input))
	reader.Comma = ' '
	args, err := reader.Read()
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil
	}
	return args
}
