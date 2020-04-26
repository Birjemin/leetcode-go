## 问题
Given an input string, reverse the string word by word.

 

Example 1:
```
Input: "the sky is blue"
Output: "blue is sky the"
```

Example 2:
```
Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
```

Example 3:
```
Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
```

Note:

- A word is defined as a sequence of non-space characters.
- Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
- You need to reduce multiple spaces between two words to a single space in the reversed string.
 

Follow up:

For C programmers, try to solve it in-place in O(1) extra space.

## 分析
翻转句子

## 最高分
```golang
func reverseWords(s string) string {
    var words []string
    var start, end int
    
    for end < len(s) {
        for start < len(s) && s[start] == ' ' {
            start++
        }
        if start == len(s) {
            break
        }
        end = start+1
        for end < len(s) && s[end] != ' ' {
            end++
        }
        words = append(words, s[start:end])
        start = end+1
    }
    
    reverse(words)
    return strings.Join(words, " ")
}

func reverse(words []string) {
    for i, j := 0, len(words)-1; i < j; i, j=i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
}
```

## 实现
最基本的实现（效率不高）
```golang
func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    s += " "

    var start int
    var ret string
    var flag bool
    for i, v := range s {
        if v != ' ' {
            flag = false
            continue
        }
        if !flag {
            ret = s[start:i] + " " + ret
            flag = true
        }
        start = i + 1
    }
    ret = strings.TrimSpace(ret)
    return ret
}
```

## 改进
效率还行
```golang
func reverseWords1(s string) string {

    // head and end space
    s = strings.TrimSpace(s)

    var tmp []string

    var start int

    // get array of vocabulary
    s += " "
    var flag bool
    var length int
    for i, v := range s {
        if v != ' ' {
            flag = false
            continue
        }
        if !flag {
            length++
            tmp = append(tmp, s[start:i])
            flag = true
        }
        start = i + 1
    }

    if len(tmp) > 0 {
        for i := 0; i < length/2; i++ {
            tmp[i], tmp[length-i-1] = tmp[length-i-1], tmp[i]
        }
    }

    return strings.Join(tmp, " ")
}

```

## 改进
将第一次解法使用buffer
```golang

```

## 反思

## 参考