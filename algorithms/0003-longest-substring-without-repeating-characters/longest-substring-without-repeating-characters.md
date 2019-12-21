## 问题

Given a string, find the length of the longest substring without repeating characters.

Example 1:
```
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## 分析
滑块思想

## 最高分
```golang
func lengthOfLongestSubstring(str string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range str {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
    dict := [128]bool{}
    length, max := 0, 0
    for i, j := 0, 0; i < len(s); i++ {
        index := s[i]
        if dict[index] {
            for ;dict[index]; j++ {
                length--
                dict[s[j]] = false
            }
        }
        dict[index] = true
        length++
        if length > max {
            max = length
        }
    }
    return max
}
```


## 实现
```golang
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
		cal = i + 1 - left
		if cal > max {
			max = cal
		}
	}
	return max
}
```

## 改进
新思路性能一般。。。。

```golang
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
```

## 反思
这个算法有点意思，可以好好深究~~~

## 参考
1. [https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/2158/Golang-solution](https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/2158/Golang-solution)
2. [https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/441906/Golang%3A-explanation-100-speed-and-memory](https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/441906/Golang%3A-explanation-\(100-speed-and-memor\))