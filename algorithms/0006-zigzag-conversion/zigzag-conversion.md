## 问题
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:
```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

Example 2:
```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
```

Explanation:
```
P     I    N
A   L S  I G
Y A   H R
P     I
```

## 分析
Z字形排列(横过来看)，实质上就是找规律的题目

## 最高分
疑问：为啥不用flag??
```golang
func convert(s string, numRows int) string {
    var direct string = "down"
    var index int = 0
    var answer string
    result := make([]string , numRows)
    if (numRows == 1){
        return s
    }else {
        for i:=0; i<len(s); i++ {
            result[index] = result[index] + string(s[i])
            if (index == (numRows-1)){
                direct = "up"
            }else if (index == 0) {
                direct = "down"
            }
            if (direct == "up") {
                index = index - 1
            }else {
                index = index + 1
            }
        }
        for i:=0; i<numRows; i++{
            answer = answer + result[i]
        }
    }
}
```


## 实现
找出规律，进行代码改进即可。

```golang
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    var str []uint8
    var flag bool
    length, num := len(s), 2*numRows-2
    // circle
    for i := 0; i < numRows; i++ {
        if i == 0 || i == numRows-1 {
            for j := i; j < length; j = j + num {
                str = append(str, s[j])
            }
        } else {
            for j := i; j < length; flag = !flag {
                str = append(str, s[j])
                if flag {
                    j = j + 2*i
                } else {
                    j = j + num - 2*i
                }
            }
        }
        flag = false
    }
    return string(str)
}
```

## 改进
不使用append(buffer刘辟！！！)
```golang
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    str := bytes.Buffer{}
    var flag bool
    length, num := len(s), 2*numRows-2
    // circle
    for i := 0; i < numRows; i++ {
        if i == 0 || i == numRows-1 {
            for j := i; j < length; j = j + num {
                str.WriteByte(s[j])
            }
        } else {
            for j := i; j < length; flag = !flag {
                str.WriteByte(s[j])
                if flag {
                    j = j + 2*i
                } else {
                    j = j + num - 2*i
                }
            }
        }
        flag = false
    }
    return str.String()
}
```

## 反思

## 参考