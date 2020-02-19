## 问题
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:
```
Input: num1 = "2", num2 = "3"
Output: "6"
```

Example 2:
```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

Note:

- The length of both num1 and num2 is < 110.
- Both num1 and num2 contain only digits 0-9.
- Both num1 and num2 do not contain any leading zero, except the number 0 itself.
- You must not use any built-in BigInteger library or convert the inputs to integer directly.

## 分析
乘法的计算，竖式法和优化竖式法求解，使用int计算时，会存在越界的情况，这就是为什么是字符串的形式，就算越界了也要结果符合预期

## 最高分
```golang
func multiply(num1 string, num2 string) string {
	num1, num2 = "0"+num1, "0"+num2 // sentinel to avoid special handling for the last carry over
	res := make([]byte, len(num1)+len(num2))
	for i, _ := range res { // initialize with byte character '0'
		res[i] = '0'
	}

	for i := 0; i < len(num2); i++ {
		pos2 := len(num2) - 1 - i
		ch2 := num2[pos2]
		co := 0
		for j := 0; j < len(num1); j++ {
			pos1 := len(num1) - 1 - j
			ch1 := num1[pos1]
			resPos := len(res) - 1 - (i + j)

			mul := int(ch1-'0')*int(ch2-'0') + int(res[resPos]-'0') + co
			res[resPos] = byte(mul%10 + '0')
			co = mul / 10
		}
	}

	var b bytes.Buffer
	firstZero := true
	for i := 0; i < len(res); i++ {
		c := res[i]
		if firstZero && c == '0' {
			continue
		}
		firstZero = false
		b.WriteByte(c)
	}

	resStr := b.String()
	if resStr == "" {
		return "0"
	}
	return resStr
}
```

## 实现
计算结果，然后计算进位，然后转变为字符串
```golang
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    var carry int
    str1, str2 := []byte(num1), []byte(num2)
    length1, length2 := len(str1), len(str2)
    temp := make([]int, length1+length2)
    // a * b
    for i := 0; i < length1; i++ {
        for j := 0; j < length2; j++ {
            temp[i+j+1] += int(str1[i]-'0') * int(str2[j]-'0')
        }
    }
    // handle carry
    for k := length1 + length2 - 1; k >= 0; k-- {
        temp[k] += carry
        temp[k], carry = temp[k]%10, temp[k]/10
    }
    // if high bit's carry is equal to 0
    if temp[0] == 0 {
        temp = temp[1:]
    }
    // buffer is awesome
    var b bytes.Buffer
    for _, v := range temp {
        b.WriteByte(byte(v + '0'))
    }
    return b.String()
}
```

## 改进
```golang

```

## 反思

## 参考
