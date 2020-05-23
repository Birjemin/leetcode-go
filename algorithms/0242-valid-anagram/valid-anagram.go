package valid_anagram

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 ,t1 := []rune(s), []rune(t)
	m := make(map[rune]int, len(s))


	for i := range t {
		m[s1[i]]++
		m[t1[i]]--
	}

	for _, b := range m {
		if b != 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	s1, t1 := str(s), str(t)
	sort.Sort(s1)
	sort.Sort(t1)
	return string(s1) == string(t1)
}

type str []rune

func (s str) Len() int {
	return len(s)
}

func (s str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s str) Less(i, j int) bool {
	return s[i] > s[j]
}