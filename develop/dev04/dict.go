package dictionary

import (
	"slices"
	"sort"
	"sync"
	"unicode"
)

func MakeDictionary(words []string) (dict map[string][]string) {
	dict = make(map[string][]string)
	keys := make(map[string]string)
	for _, word := range words {
		parseWord(word, dict, keys)
	}
	wg := &sync.WaitGroup{}
	for key := range dict {
		wg.Add(1)
		go sortWords(dict[key], wg)
	}
	wg.Wait()
	return
}

func parseWord(word string, dict map[string][]string, keys map[string]string) {
	chars := []rune(word)
	for i := range chars {
		chars[i] = unicode.ToLower(chars[i])
	}
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

func sortWords(words []string, wg *sync.WaitGroup) {
	defer wg.Done()
	slices.Sort(words)
}
