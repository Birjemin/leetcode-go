## 问题
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example 2:
```
Input: "cbbd"
Output: "bb"
```

## 分析
输出最大的回文字符串

## 最高分

```golang
func longestPalindrome(s string) string {
	min := 0
	max := 0
	slen := len(s)

	for i := 0; i < slen; i++ {
	    // find longest palindromic substring with odd length and center at i
		l := i
		h := i
		for l >= 0 && h < slen && s[l] == s[h] {
			l--
			h++
		}
		if h-l-1 > max-min {
			min = l + 1
			max = h
		}

	    // find longest palindromic substring with even length and center at i,i+1
		l = i
		h = i + 1
		for l >= 0 && h < slen && s[l] == s[h] {
			l--
			h++
		}
		if h-l-1 > max-min {
			min = l + 1
			max = h
		}
	}

	return s[min:max]
}
```

## 实现
最基本的实现方式
```golang
func longestPalindrome(s string) string {
    var res string
    var temp string
    length := len(s)
    var l, l2 int
    if length <= 1 {
        return s
    }
    for i := 0; i < length-1; i++ {
        for j := i+l; j <= length; j++ {
            temp = s[i:j]
            l2 = len(temp)
            if isPalindrome(temp, l2)  {
                res = temp
                l = l2
            }
        }
    }
    return res
}

// 判断某一个是否是回文
func isPalindrome(s string, j int) bool {
    i:= 0
    for i < j {
        if s[i] != s[j-1] {
            return false
        }
        i++
        j--
    }
    return true
}
```

## 改进
根据最高分修改，分偶数和奇数
```golang
func longestPalindrome(s string) string {
    min, max, length := 0, 0, len(s)
    for i := 0; i < length; i++ {
        // 奇数
        l, h := i, i
        for l >= 0 && h < length && s[l] == s[h] {
            l--
            h++
        }
        if h-l-1 > max-min {
            min, max = l+1, h
        }

        // 偶数
        l, h = i, i+1
        for l >= 0 && h < length && s[l] == s[h] {
            l--
            h++
        }
        if h-l-1 > max-min {
            min, max = l+1, h
        }
    }
    return s[min:max]
}
```

## 反思

## 参考