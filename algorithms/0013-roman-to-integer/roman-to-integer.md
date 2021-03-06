## 问题
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

```
Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
Example 3:

Input: "IX"
Output: 9
Example 4:

Input: "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 5:

Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

## 分析
从后往前计算，如果前面一位小于当前位，则是减法，反之加法，比如MCMXCIV = 5 - 1 + 100 - 10 + 1000 - 100 + 1000

## 最高分

```golang
func romanToInt(s string) int {
    n := len(s)
    if n < 1 {
        return 0
    }
    table := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    val := 0
    for i := 0; i < n - 1; i++ {
        a := rune(s[i])
        b := rune(s[i + 1])
        if table[a] < table[b] {
            val -= table[a]
        } else {
            val += table[a]
        }
    }
    return val + table[rune(s[n - 1])]
}
```

## 实现
最基本的实现，就是循环判断
```golang
func romanToInt(s string) int {
    table := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    length := len(s)
    sum := table[rune(s[length - 1])]
    for i := length - 1; i > 0; i-- {
        z := table[rune(s[i-1])]
        if table[rune(s[i])] > z {
            sum -= z
        } else {
            sum += z
        }
    }
    return sum
}
```

## 改进
每次循环保存上一次值，不在重新获取
```golang

```
* 时间复杂度：O(n)
* 空间复杂度：O(n)

## 反思

## 参考