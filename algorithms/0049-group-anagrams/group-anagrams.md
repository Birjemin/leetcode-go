## 问题
Given an array of strings, group anagrams together.

Example:
```
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

Note:

All inputs will be in lowercase.
The order of your output does not matter.

## 分析
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

## 最高分
直接使用一个map来存储，然后转变为[][]string类型
```golang
func groupAnagrams(strs []string) [][]string {
    hm := make(map[string][]string)
    for _, w := range strs {
        key := sortString(w)
        hm[key] = append(hm[key], w)
    }
    
    var ans [][]string
    for _, words := range hm {
        ans = append(ans, words)
    }
    return ans
}

// important!!!
func sortString(str string) string {
    runes := []rune(str)
    sort.Slice(runes, func(a, b int) bool {
        return runes[a] < runes[b]
    })
    return string(runes)
}
```

## 实现
先排序，然后找到相同的数值，实现Sort接口
```golang
type bytes []byte

func (b bytes) Len() int {
    return len(b)
}

func (b bytes) Less(i, j int) bool {
    return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
    var ret [][]string
    var j int
    temp := make(map[string]int)
    length := len(strs)
    for i := 0; i < length; i++ {
        key := bytes(strs[i])
        sort.Sort(key)
        str := string(key)
        // exist
        if _, ok := temp[str]; ok {
            ret[temp[str]] = append(ret[temp[str]], strs[i])
        } else {
            temp[str] = j
            ret = append(ret, []string{strs[i]})
            j++
        }
    }
    return ret
}
```

## 改进
```golang

```

## 反思

## 参考