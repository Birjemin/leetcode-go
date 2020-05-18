package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	// s -> t t -> s
	m1, m2 := make(map[rune]rune, len(s)), make(map[rune]rune, len(t))

	for i, b := range s {
		tmp := rune(t[i])
		if v, ok := m1[b]; ok {
			// had existed
			if tmp != v {
				return false
			}
		} else {
			if _, ok := m2[tmp]; ok {
				return false
			}
			m1[b], m2[tmp] = tmp, b
		}
	}
	return true
}

func isIsomorphic1(s string, t string) bool {
	// s -> t t -> s
	m1, m2 := make(map[uint8]int, len(s)), make(map[uint8]int, len(t))
	for idx := range s {
		if m1[s[idx]] != m2[t[idx]] {
			return false
		}
		// Why idx+1?? map[uint8]int default 0 if the value is not existed
		m1[s[idx]], m2[t[idx]] = idx+1, idx+1
	}
	return true
}
