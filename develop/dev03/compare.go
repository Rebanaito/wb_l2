package main

import (
	"strings"
	"unicode"
)

func compare(left, right string, flags options) bool {
	var result bool
	if flags.k != 0 {
		result = substring(left, right, flags)
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

func substring(l, r string, flags options) bool {
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

func parseInt(str string) (num int) {
	chars := []rune(str)
	i := 0
	for i < len(chars) && unicode.IsSpace(chars[i]) {
		i++
	}
	negative := false
	if i < len(chars) && chars[i] == '-' {
		negative = true
		i++
	} else if i < len(chars) && chars[i] == '+' {
		i++
	}
	for i < len(chars) && chars[i] >= '0' && chars[i] <= '9' {
		num = num*10 + int(chars[i]-'0')
		i++
	}
	if negative {
		num = -num
	}
	return
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
