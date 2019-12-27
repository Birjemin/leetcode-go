## 问题
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

```
Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
```

## 分析
比较简单

## 最高分
```golang

```


## 实现
```golang
func longestCommonPrefix(strs []string) string {
    length := len(strs)
    if length == 0 {
        return ""
    } else if length == 1 {
        return strs[0]
    }

    var flag string
    for i, s := range strs[0] {
        for j := 1; j < length; j++ {
            if len(strs[j]) <= i || strs[j][i] != byte(s) {
                return flag
            }
        }
        flag += string(s)
    }
    return flag
}
```

## 改进
减少每次判断长度
```golang
func longestCommonPrefix(strs []string) string {
    length := len(strs)
    if length == 0 {
        return ""
    } else if length == 1 {
        return strs[0]
    }

    var flag []byte
    lenMap := make(map[int]int, length - 1)
    for i, s := range strs[0] {
        for j := 1; j < length; j++ {
            if lenMap[j] == 0 {
                lenMap[j] = len(strs[j])
            }
            if lenMap[j] <= i || strs[j][i] != byte(s) {
                return string(flag)
            }
        }
        flag = append(flag, byte(s))
    }
    return string(flag)
}
```

## 反思
无

## 参考