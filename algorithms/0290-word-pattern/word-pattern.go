package word_pattern

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")

	len1, len2 := len(pattern), len(s)
	if len1 != len2 {
		return false
	}

	kv := make(map[uint8]string, len1)
	for i, v := range s {
		if j, ok := kv[pattern[i]]; ok {
			if j != v {
				return false
			}
		} else {
			kv[pattern[i]] = v
		}
	}

	kv1 := make(map[string]rune, len2)
	for i, v := range pattern {
		if j, ok := kv1[s[i]]; ok {
			if j != v {
				return false
			}
		} else {
			kv1[s[i]] = v
		}
	}

	return true
}

func wordPattern1(pattern string, str string) bool {
	s := strings.Split(str, " ")

	len1, len2 := len(pattern), len(s)
	if len1 != len2 {
		return false
	}

	kv, kv1 := make(map[uint8]string, len1), make(map[string]uint8, len2)

	for i := 0; i < len1; i++ {
		if v, ok := kv[pattern[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			kv[pattern[i]] = s[i]
		}

		if v, ok := kv1[s[i]]; ok {
			if v != pattern[i] {
				return false
			}
		} else {
			kv1[s[i]] = pattern[i]
		}
	}

	return true
}
