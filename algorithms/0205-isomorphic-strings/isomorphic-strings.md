## 问题
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:
```
Input: s = "egg", t = "add"
Output: true
```

Example 2:
```
Input: s = "foo", t = "bar"
Output: false
```

Example 3:
```
Input: s = "paper", t = "title"
Output: true
```

Note:
You may assume both s and t have the same length.

## 分析
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

map映射处理就好了

## 最高分
```golang
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}

	return true
}
```

## 实现
建立映射关系解决最简单
```golang
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
```

## 改进
使用映射关系，将`s[i]`和`t[i]`的值映射为同一个值
```golang
func isIsomorphic1(s string, t string) bool {
	// s -> t t -> s
	m1, m2 := make(map[uint8]int, len(s)), make(map[uint8]int, len(t))
	for idx := range s {
		if m1[s[idx]] != m2[t[idx]] {
			return false
		}
		m1[s[idx]], m2[t[idx]] = idx+1, idx+1
	}
	return true
}
```

## 反思

## 参考