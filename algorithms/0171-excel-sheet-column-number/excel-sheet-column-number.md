## 问题
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:
```
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
    ...
```

Example 1:
```
Input: "A"
Output: 1
```

Example 2:
```
Input: "AB"
Output: 28
```

Example 3:
```
Input: "ZY"
Output: 701
```

## 分析
和168题类似，只是它的逆过程

## 最高分
```golang
func titleToNumber(s string) int {
    n, p, l := 0, 1, len(s)-1
    for i := l; i >= 0; i-- {
        n += p * (int(s[i]-'A') + 1)
        p *= 26
    }
    return n
}
```


## 实现
```golang
func titleToNumber(s string) int {
    rate, sum := 1, 0
    for i := len(s) - 1; i >= 0; i-- {
        //sum += rate * (int(s[i]) - 64)
        sum += rate * (int(s[i]-'A'+1))
        rate *= 26
    }
    return sum
}
```

## 改进
int(v-'A'+1) 效率比int(v-64)高。。
```golang
func titleToNumber(s string) int {
    var sum int
    for _, v := range s {
        sum = 26*sum + int(v-'A'+1)
    }
    return sum
}
```

## 反思

## 参考