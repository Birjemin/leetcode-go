## 问题
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
```
Input: haystack = "hello", needle = "ll"
Output: 2
```
Example 2:
```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

## 分析
可以使用kmp算法（这个比较有趣）

## 最高分
```golang

```


## 实现
内置函数方法？？
```golang
func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}
```

## 改进
切割就好
```golang
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    length, l := len(haystack), len(needle)
    for idx := 0; idx <= length-l; idx++ {
        if haystack[idx:l+idx] == needle {
            return idx
        }
    }
    return -1
}
```

## 反思

## 参考
