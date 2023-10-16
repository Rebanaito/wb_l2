package dictionary

import (
	"sort"
)

func MakeDictionary(words []string) map[string][]string {
	dict := make(map[string][]string)
	keys := make(map[string]string)
	for _, word := range words {
		chars := []rune(word)
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		value, ok := keys[string(chars)]
		if !ok {
			keys[string(chars)] = word
			dict[word] = append(dict[word], word)
		} else {
			dict[value] = append(dict[value], word)
		}
	}
	return dict
}
