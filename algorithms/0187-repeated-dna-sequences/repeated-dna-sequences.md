## 问题
All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example:
```
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Output: ["AAAAACCCCC", "CCCCCAAAAA"]
```

## 分析
编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。
和136题很像，不知道可不可以使用136的位运算
## 最高分
```golang
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	charMap, mp, result := map[uint8]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}, make(map[uint32]int, 0), []string{}
	var cur uint32
	for i := 0; i < 9; i++ { // 前9位，忽略
		cur = cur<<2 | charMap[s[i]]
	}
	for i := 9; i < len(s); i++ {
		cur = ((cur << 2) & 0xFFFFF) | charMap[s[i]]
		if mp[cur] == 0 {
			mp[cur] = 1
		} else if mp[cur] == 1 { // >2，重复
			mp[cur] = 2
			result = append(result, s[i-9:i+1])
		}
	}
	return result
}
```

## 实现
最简单的方式
```golang
func findRepeatedDnaSequences(s string) []string {
	var ret []string
	tmp := make(map[string]bool)

	for i := len(s) - 10; i >= 0; i-- {
		t := s[i : i+10]
		if _, ok := tmp[t]; ok {
			if tmp[t] {
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
```

## 改进
改进Rabin-Karp 算法
```golang
func findRepeatedDnaSequences(s string) []string {
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
```

## 反思

## 参考