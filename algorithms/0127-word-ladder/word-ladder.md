## 问题
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

* Only one letter can be changed at a time.
* Each transformed word must exist in the word list. Note that beginWord is not a transformed word.

Note:

- Return 0 if there is no such transformation sequence.
- All words have the same length.
- All words contain only lowercase alphabetic characters.
- You may assume no duplicates in the word list.
- You may assume beginWord and endWord are non-empty and are not the same.

Example 1:
```
Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
```

Example 2:
```
Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
```

## 分析
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。（题目看懂了，但是没有思路）

- 感觉属于高级难度，暂时放着不做
## 最高分
```golang

```

## 实现
bfs超时，看官方解答，问题出现于dislocate方法
```golang
func ladderLength(beginWord string, endWord string, wordList []string) int {
    return bfs([]string{beginWord}, endWord, 1, [][]string{wordList})
}

func bfs(beginWords []string, endWord string, count int, wordLists [][]string) int {
    // cal
    var queues []string
    var extra [][]string

    if len(beginWords) == 0 {
        return 0
    }
    for i, b := range beginWords {
        if endWord == b {
            return count
        }
        for j, w := range wordLists[i] {
            if !dislocate(b, w) {
                continue
            }
            // find
            queues = append(queues, w)
            wordLis := delElement(j, wordLists[i])
            extra = append(extra, wordLis)
        }
    }
    return bfs(queues, endWord, count+1, extra)
}

// delete element
func delElement(idx int, wordList []string) []string {
    length := len(wordList)
    b := make([]string, length)
    copy(b, wordList)
    copy(b[idx:], b[idx+1:])
    return b[:length-1]
}

// find the element
func dislocate(s1, s2 string) bool {
    var count int
    for i, s := range s1 {
        if uint8(s) != s2[i] {
            count++
        }
        if count > 1 {
            return false
        }
    }
    if count == 1 {
        return true
    }
    return false
}

```

## 改进
bfs得到一条路径，将这条路径优化成最短路径即可
```golang
func ladderLength(beginWord string, endWord string, wordList []string) int {
    return bfs([]string{beginWord}, []string{""}, wordList, endWord, 1, len(wordList))
}

func bfs(beginWords, beforeWords, wordList []string, endWord string, count, max int) int {
    // cal
    var queues []string
    var beforeQueues []string

    if max < count {
        return 0
    }

    for i, b := range beginWords {
        if b == endWord {
            return count
        }
        for _, w := range wordList {
            // can not turn back
            if w == beforeWords[i] {
                continue
            }
            // find the next element
            if !dislocate(b, w) {
                continue
            }
            // find
            queues = append(queues, w)
            beforeQueues = append(beforeQueues, b)
        }
    }
    return bfs(queues, beforeQueues, wordList, endWord, count+1, max)
}

// find the element
func dislocate(s1, s2 string) bool {
    var count int
    for i, s := range s1 {
        if uint8(s) != s2[i] {
            count++
        }
        if count > 1 {
            return false
        }
    }
    if count == 1 {
        return true
    }
    return false
}
```

## 反思

## 参考