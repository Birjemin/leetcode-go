## 问题
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
```
Input: s = "anagram", t = "nagaram"
Output: true
```

Example 2:
```
Input: s = "rat", t = "car"
Output: false
```

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

## 分析
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

和49题类似

字母序排序，然后判断是否相等即可

## 最高分
其实就是一个减，一个加
```golang
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
```

## 实现
最简单的实现，先排序，后判断是否相等
```golang
func isAnagram(s string, t string) bool {
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
```

## 改进
使用一个map来保存即可
```golang
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, len(s))
	for _, b := range s {
		m[b]++
	}

	for _, b := range t {
		if v, ok := m[b]; !ok {
			return false
		} else if v == 0 {
			return false
		} else {
			m[b]--
		}
	}
	return true
}
```

## 反思

## 参考