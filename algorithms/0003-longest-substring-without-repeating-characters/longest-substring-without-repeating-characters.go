package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var freq [256]int
	cal, max, left := 0, 0, 0
	length := len(s)
	for i := 0; i < length; i++ {
		// 窗口里面出现过
		if freq[s[i]] > left {
			left = freq[s[i]]
		}
		// 该字符串出现的位置
		freq[s[i]] = i + 1
		cal = freq[s[i]] - left
		if cal > max {
			max = cal
		}
	}
	return max
}

func lengthOfLongestSubstring1(s string) int {
	m, max, left, cal := make(map[rune]int), 0, 0, 0
	// 序号和值
	for idx, c := range s {
		if _, okay := m[c]; okay == true && m[c] > left {
			left = m[c]
		}
		m[c] = idx + 1
		cal = m[c] - left
		if cal > max {
			max = cal
		}
	}
	return max
}
