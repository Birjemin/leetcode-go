## 问题
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:
```
Input: pattern = "abba", str = "dog cat cat dog"
Output: true
```

Example 2:
```
Input:pattern = "abba", str = "dog cat cat fish"
Output: false
```

Example 3:
```
Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
```

Example 4:
```
Input: pattern = "abba", str = "dog dog dog dog"
Output: false
```

Notes:

You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.

## 分析
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

没想到好的办法，两个map解决

## 最高分
```golang
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) == 0 {
		return false
	}
	if len(strs) != len(pattern) {
		return false
	}

	appearance := make(map[string]bool)
	mapping := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		ch, str := pattern[i], strs[i]

		mapped, ok := mapping[ch]
		if ok {
			if mapped == str {
				continue
			}
			return false
		}

		appeared, ok := appearance[str]
		if ok && appeared {
			return false
		}

		mapping[ch] = str
		appearance[str] = true
	}
	return true
}
```

## 实现
最直接的方式
```golang
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
```

## 改进
优化空间复杂度
```golang
func wordPattern(pattern string, str string) bool {
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

```

## 反思

## 参考