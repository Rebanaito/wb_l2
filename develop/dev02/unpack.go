package unpack

import (
	"errors"
	"strings"
)

var InvalidString error = errors.New("Invalid input string")

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	var unpacked strings.Builder
	chars := []rune(input)
	i := 0
	for i < len(chars) {
		if i == len(chars)-1 && chars[i] == '\\' {
			return "", InvalidString
		} else if chars[i] == '\\' {
			escape(chars, &i, &unpacked)
		} else if isDigit(chars[i]) {
			return "", InvalidString
		} else {
			checkNext(chars, &i, &unpacked)
		}
	}
	if unpacked.Len() == 0 {
		return "", InvalidString
	}
	return unpacked.String(), nil
}

func checkNext(chars []rune, i *int, unpacked *strings.Builder) {
	if *i < len(chars)-1 && isDigit(chars[*i+1]) {
		len := int(chars[*i+1] - '0')
		for j := 0; j < len; j++ {
			unpacked.WriteRune(chars[*i])
		}
		*i += 2
	} else {
		unpacked.WriteRune(chars[*i])
		*i += 1
	}
}

func escape(chars []rune, i *int, unpacked *strings.Builder) {
	char := chars[*i+1]
	if *i < len(chars)-2 && isDigit(chars[*i+2]) {
		len := int(chars[*i+2] - '0')
		for j := 0; j < len; j++ {
			unpacked.WriteRune(char)
		}
		*i += 3
	} else {
		unpacked.WriteRune(char)
		*i += 2
	}
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
