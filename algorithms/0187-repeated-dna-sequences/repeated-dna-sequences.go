package repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
	var ret []string
	tmp := make(map[string]bool)

	for i := len(s) - 10; i >= 0; i-- {
		t := s[i : i+10]
		if v, ok := tmp[t]; ok {
			if v {
				continue
			}
			ret = append(ret, t)
			tmp[t] = true
		} else {
			tmp[t] = false
		}
	}

	return ret
}

func findRepeatedDnaSequences1(s string) []string {
	if len(s) < 10 {
		return nil
	}
	var result []string
	var cur uint32

	charMap, mp := map[uint8]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}, make(map[uint32]int, 0)

	for i := 0; i < 9; i++ {
		cur = cur<<2 | charMap[s[i]]
	}
	for i := 9; i < len(s); i++ {
		cur = ((cur << 2) & 0xFFFFF) | charMap[s[i]]
		if mp[cur] == 0 {
			mp[cur] = 1
		} else if mp[cur] == 1 { //
			mp[cur] = 2
			result = append(result, s[i-9:i+1])
		}
	}
	return result
}
