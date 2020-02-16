## 问题
The count-and-say sequence is the sequence of integers with the first five terms as following:
```
1.     1
2.     11
3.     21
4.     1211
5.     111221
```
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence. You can do so recursively, in other words from the previous member read off the digits, counting the number of digits in groups of the same digit.

Note: Each term of the sequence of integers will be represented as a string.

 

Example 1:
```
Input: 1
Output: "1"
Explanation: This is the base case.
```

Example 2:
```
Input: 4
Output: "1211"
Explanation: For n = 3 the term was "21" in which we have two groups "2" and "1", "2" can be read as "12" which means frequency = 1 and value = 2, the same way "1" is read as "11", so the answer is the concatenation of "12" and "11" which is "1211".
```

## 分析
外观数列，开始不理解是个啥，然后网上找了找资料才知道这是个啥。。。
其实就是当前行对上一行的描述，比如第5行描述第四行1个1，1个2，2个1，所以：111221

## 最高分
```golang

```


## 实现
最简单的就是递归
```golang
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    } else {
        str := countAndSay(n - 1)
        var flag rune
        var res string
        var count int
        for _, s := range str {
            if flag == 0 {
                flag = s
            }
            if flag == s {
                count++
            } else {
                res += strconv.Itoa(count) + string(flag)
                flag = s
                count = 1
            }
        }
        res += strconv.Itoa(count) + string(flag)
        return res
    }
}
```

## 改进
使用strings.Builder
```golang
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    } else {
        str := countAndSay(n - 1)
        var flag rune
        var count int
        var res strings.Builder
        for _, s := range str {
            if flag == 0 {
                flag = s
            }
            if flag == s {
                count++
            } else {
                res.WriteString(strconv.Itoa(count))
                res.WriteRune(flag)
                flag = s
                count = 1
            }
        }
        res.WriteString(strconv.Itoa(count))
        res.WriteRune(flag)
        return res.String()
    }
}
```

## 反思

## 参考
