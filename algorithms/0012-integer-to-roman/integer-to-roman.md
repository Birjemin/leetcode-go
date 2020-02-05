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
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:
```
Input: 3
Output: "III"
```
Example 2:
```
Input: 4
Output: "IV"
```
Example 3:
```
Input: 9
Output: "IX"
```
Example 4:
```
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```
Example 5:
```
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

## 分析
和13题类似类似，是从阿拉伯数字到罗马数字，分解因子

## 最高分
实现很简练，比第二种快
```golang
func intToRoman(num int) string {
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	i := 0
	for num > 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		result += romans[i]
	}
	return result
}

func intToRoman(num int) string {
    d := [4][]string{
        {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
        {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
        {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
        {"", "M", "MM", "MMM"},
    }
    return d[3][num/1000] + d[2][num/100%10] + d[1][num/10%10] + d[0][num%10]
}
```


## 实现
借鉴了其他方案的实现，这个有点想不到。。。(改用Buffer)
```golang
func intToRoman(num int) string {
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    str := bytes.Buffer{}
    // cal
    for i := 0; num > 0; num -= values[i] {
        for values[i] > num {
            i++
        }
        str.WriteString(romans[i])
    }
    return str.String()
}
```

## 改进
```golang

```

## 反思

## 参考
