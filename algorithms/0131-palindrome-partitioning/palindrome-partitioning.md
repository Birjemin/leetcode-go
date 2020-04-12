## 问题
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:
```
Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
```

## 分析
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
- dfs遍历，这个很高阶了~~~(想了老一会儿)

## 最高分
```golang
func partition(s string) [][]string {
    result := [][]string{}
    recPart(s, []string{}, &result)
    return result
}

func recPart(s string, c []string, result *[][]string) {
    fmt.Println(s)
    if len(s) == 0 {
        *result = append(*result, c)
    }
    for i := 1; i <= len(s); i++ {
        subS := s[:i]
        if isPal(subS) {
            newC := make([]string, len(c)+1)
            copy(newC, c)
            newC[len(newC)-1] = subS
            recPart(s[i:], newC, result)
        }
    }
}

func isPal(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
```

## 实现
dfs实现，需要想一会儿，涉及到拷贝和循环遍历
```golang
func partition(s string) [][]string {
    var ret [][]string
    dfs([]string{}, s, &ret)
    return ret
}

func dfs(prefix []string, s string, ret *[][]string) {
    if s == "" {
        b := make([]string, len(prefix))
        copy(b, prefix)
        *ret = append(*ret, b)
        return
    }

    for i := range s {
        if isPalindrome(i, s) {
            dfs(append(prefix, s[:i+1]), s[i+1:], ret)
        }
    }
}

// palindrome
func isPalindrome(end int, s string) bool {
    for i := 0; i < end; i, end = i+1, end-1 {
        if s[i] != s[end] {
            return false
        }
    }
    return true
}
```

## 改进
```golang

```

## 反思

## 参考