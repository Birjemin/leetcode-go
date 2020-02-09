## 问题
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
```
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
```
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

## 分析
遍历得到所有的可能性，考察的是dfs搜索

## 最高分
```golang

```


## 实现
递归遍历，这么实现很简单
```golang
var (
    m = [][]string{
        {"a", "b", "c"},
        {"d", "e", "f"},
        {"g", "h", "i"},
        {"j", "k", "l"},
        {"m", "n", "o"},
        {"p", "q", "r", "s"},
        {"t", "u", "v"},
        {"w", "x", "y", "z"},
    }
)

func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    return find(digits, []string{})
}

func find(str string, s []string) []string {
    var ret []string
    if str == "" {
        return s
    }
    length := len(s)
    for _, v := range m[str[0]-'0'-2] {
        if length == 0 {
            ret = append(ret, string(v))
        } else {
            for _, o := range s {
                ret = append(ret, string(o)+string(v))
            }
        }
    }
    return find(str[1:], ret)
}
```

## 改进
```golang

```

## 反思

## 参考
