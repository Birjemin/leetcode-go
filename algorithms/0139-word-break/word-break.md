## 问题
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.

Example 1:
```
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
```

Example 2:
```
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
```

Example 3:
```
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
```

## 分析
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

## 最高分
不容易理解
```golang
func wordBreak(s string, wordDict []string) bool {
    m := make(map[string]bool)
    list := make([]bool, len(s) + 1)
    for _ , v := range wordDict {
        m[v] = true
    }
    list[0] = true
    l := len(s)
    for i := 1; i <= l; i++ {
        for j := 0; j < i; j++ {
            if list[j] && m[s[j:i]] {
                list[i] = true
                break
            }
        }
    }
    return list[l]
}
```

## 实现
超时实现，得想起他方式(最后一个test示例真的是刘辟)
```golang
func wordBreak(s string, wordDict []string) bool {
    return dfs(s, len(s), wordDict)
}

func dfs(s string, sLen int, wordDict []string) bool {
    if s == "" {
        return true
    }
    for _, v := range wordDict {
        length := len(v)
        if sLen < length {
            continue
        }
        if s[:length] == v {
            if dfs(s[length:], sLen-length, wordDict) {
                return true
            }
        }
    }
    return false
}
```

## 改进
稍微改进了一点，不会出现超时，但是性能并不是很高
```golang
func wordBreak(s string, wordDict []string) bool {
    wordDictMap := make(map[string]int)
    for _, s := range wordDict {
        wordDictMap[s] = len(s)
    }
    return bfs(map[string]int{s: len(s)}, wordDictMap)
}

func bfs(str map[string]int, wordDictMap map[string]int) bool {
    if len(str) == 0 {
        return false
    }
    queue := make(map[string]int)
    for s, sLen := range str {
        for w, wLen := range wordDictMap {
            if sLen < wLen {
                continue
            }
            if s[:wLen] == w {
                left := s[wLen:]
                if left == "" {
                    return true
                }
                // if not exist in queue
                if _, ok := queue[left]; !ok {
                    queue[left] = sLen - wLen
                }
            }
        }
    }
    return bfs(queue, wordDictMap)
}
```

## 改进
继续改进，比如每一次迭代queue都是新生成的，如果之前算过，这一次还要计算，可是这个是没有必要的，可以简化(可以通过打日志发现之间的联系)
```golang
func wordBreak(s string, wordDict []string) bool {
    // got length of wordDict and record memo
    wordDictMap, memo := make(map[string]int), make(map[string]int)
    for _, s := range wordDict {
        wordDictMap[s] = len(s)
    }

    return bfs(map[string]int{s: len(s)}, wordDictMap, &memo)
}

func bfs(str map[string]int, wordDictMap map[string]int, memo *map[string]int) bool {
    if len(str) == 0 {
        return false
    }
    queue := make(map[string]int)
    for s, sLen := range str {
        for w, wLen := range wordDictMap {

            // not fit or got it
            if sLen < wLen || s[:wLen] != w {
                continue
            }

            left := s[wLen:]

            // found!  end
            if left == "" {
                return true
            }

            // if it has in memo, just continue and it is not necessary to be stored in queue
            if _, ok := (*memo)[left]; ok {
                continue
            }

            (*memo)[left] = sLen - wLen

            // if not exist in queue, stored the value
            if _, ok := queue[left]; !ok {
                queue[left] = sLen - wLen
            }
        }
    }
    return bfs(queue, wordDictMap, memo)
}
```

## 改进
使用dp回溯算法（和上面使用memo标记类似的，需要理解的是s[j:i]的含义，dp[i]标识0~i可以被wordDict中字符串组合！！！）
```golang
func wordBreak(s string, wordDict []string) bool {

    // got length of wordDict and record memo
    length := len(s)
    wordDictMap, dp := make(map[string]int, length), make(map[int]bool, length+1)
    for _, s := range wordDict {
        wordDictMap[s] = len(s)
    }
    dp[0] = true
    for i := 1; i <= length; i++ {
        for j := 0; j < i; j++ {
            if !dp[j] {
                continue
            }
            if _, ok := wordDictMap[s[j:i]]; ok {
                dp[i] = true
                break
            }
        }
    }
    return dp[length]
}
```

## 反思

## 参考