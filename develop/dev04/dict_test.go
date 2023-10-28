package dictionary

import "testing"

func TestDict(t *testing.T) {
	words := []string{"столик", "тяпка", "слиток", "листок", "пятка", "пятак"}
	dict := MakeDictionary(words)
	if len(dict) != 2 {
		t.Fatal(`Expected to have 2 subsets, have`, len(dict))
	}
	if dict["тяпка"][0] != "пятак" || dict["тяпка"][1] != "пятка" ||
		dict["тяпка"][2] != "тяпка" {
		t.Fatal(`Anagram array is not sorted properly`)
	}
	if dict["столик"][0] != "листок" || dict["столик"][1] != "слиток" ||
		dict["столик"][2] != "столик" {
		t.Fatal(`Anagram array is not sorted properly`)
	}
}
