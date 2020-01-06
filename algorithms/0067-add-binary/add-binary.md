## 问题
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:
```
Input: a = "11", b = "1"
Output: "100"
```

Example 2:
```
Input: a = "1010", b = "1011"
Output: "10101"
```

## 分析
二进制加减

## 最高分
```golang
func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }
    ans := make([]byte, len(a) + 1)
	var carry byte
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i - 1, j - 1
		var a2Digit, b2Digit byte
		if (i >= 0) {
			a2Digit = a[i] - '0'
		}
		if (j >= 0) {
			b2Digit = b[j] - '0'
		}
        // sum and carry of full adder
		sum := a2Digit ^ b2Digit ^ carry
		carry = a2Digit & b2Digit | carry & (a2Digit ^ b2Digit)
        ans[i+1] = sum + '0'
	}
	if carry == 1 {
		ans[0] = '1'
        return string(ans)
	}
    return string(ans[1:])
}
```


## 实现
最直接的实现方式
```golang
func addBinary(a string, b string) string {
    if a == "" && b == "" {
        return "0"
    }
    l1, l2 := len(a), len(b)
    // 先保证a的长度大于b的长度
    if l1 < l2 {
        a, b = b, a
        l1, l2 = l2, l1
    }
    var y, sum int
    ret := make([]byte, l1+1)
    for i := 0; i < l1; i++ {
        if i < l2 {
            y += byte2Int(b[l2-i-1])
        }
        sum = byte2Int(a[l1-1-i]) + y
        if sum == 1 {
            y = 0
            ret[l1-i] = 49
        } else if sum == 2 {
            y = 1
            ret[l1-i] = 48
        } else if sum == 3 {
            y = 1
            ret[l1-i] = 49
        } else {
            y = 0
            ret[l1-i] = 48
        }
    }
    if y == 1 {
        ret[0] = 49
    } else {
        ret = ret[1:]
    }
    return string(ret)
}

func byte2Int(s byte) int {
    if s == 49 {
        return 1
    }
    return 0
}
```

## 改进
简单的优化，采用位运算
```golang

func byte2Int(s byte) int {
    if s == 49 {
        return 1
    }
    return 0
}

func addBinary(a string, b string) string {
    if a == "" && b == "" {
        return "0"
    }
    l1, l2 := len(a), len(b)
    // 先保证a的长度大于b的长度
    if l1 < l2 {
        a, b = b, a
        l1, l2 = l2, l1
    }
    var flag, x, y int
    ret := make([]byte, l1+1)
    for i := 0; i < l1; i++ {
        if i < l2 {
            x = byte2Int(b[l2-i-1])
        } else {
            x = 0
        }
        y = byte2Int(a[l1-i-1])
        // 两个不相等，则会改变当前位
        ret[l1-i] = byte(x ^ y ^ flag + 48)
        // 两个相等，则会产生移位
        flag = x&y | flag&(x^y)
    }
    if flag == 1 {
        ret[0] = 49
    } else {
        ret = ret[1:]
    }
    return string(ret)
}
```

## 反思
位运算的使用

## 参考
