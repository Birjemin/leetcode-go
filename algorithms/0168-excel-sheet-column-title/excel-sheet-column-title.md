## 问题
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:
```
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...
```

Example 1:
```
Input: 1
Output: "A"
```

Example 2:
```
Input: 28
Output: "AB"
```

Example 3:
```
Input: 701
Output: "ZY"
```

## 分析
简单，只是简单的除法运算

## 最高分
```golang
func convertToTitle(n int) string {
    result := ""
    for n != 0 {
        result = string('A' + (n-1) % 26) + result
        n = (n-1)/26
    }
    return result
}
```


## 实现
其实就是简单的禁止转换
```golang
func convertToTitle(n int) string {
    ret := new(bytes.Buffer)
    for n > 26 {
        remainder := n % 26
        n /= 26
        if remainder == 0 {
            ret.WriteRune(rune(90))
            n -= 1
        } else {
            ret.WriteRune(rune(remainder + 64))
        }
    }
    ret.WriteRune(rune(n + 64))
    bytesArr := ret.Bytes()
    for from, to := 0, len(bytesArr)-1; from < to; from, to = from+1, to-1 {
        bytesArr[from], bytesArr[to] = bytesArr[to], bytesArr[from]
    }
    return string(bytesArr)
}
```

## 改进
```golang

```

## 反思

## 参考