## 问题
Compare two version numbers version1 and version2.
If `version1 > version2` return 1; if `version1 < version2` return -1;otherwise return 0.

You may assume that the version strings are non-empty and contain only digits and the . character.

The . character does not represent a decimal point and is used to separate number sequences.

For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

You may assume the default revision number for each level of a version number to be 0. For example, version number 3.4 has a revision number of 3 and 4 for its first and second level revision number. Its third and fourth level revision number are both 0.

 

Example 1:
```
Input: version1 = "0.1", version2 = "1.1"
Output: -1
```

Example 2:
```
Input: version1 = "1.0.1", version2 = "1"
Output: 1
```

Example 3:
```
Input: version1 = "7.5.2.4", version2 = "7.5.3"
Output: -1
```

Example 4:
```
Input: version1 = "1.01", version2 = "1.001"
Output: 0
Explanation: Ignoring leading zeroes, both “01” and “001" represent the same number “1”
```

Example 5:
```
Input: version1 = "1.0", version2 = "1.0.0"
Output: 0
Explanation: The first version number does not have a third level revision number, which means its third level revision number is default to "0"
```

Note:

- Version strings are composed of numeric strings separated by dots . and this numeric strings may have leading zeroes.
- Version strings do not start or end with dots, and they will not be two consecutive dots.

## 分析
判断版本号大小，其实就是遍历查找然后对比大小

## 最高分
```golang
import (
"strings"
"strconv"
)

func compareVersion(version1 string, version2 string) int {

    ver1Arr := strings.Split(version1, ".")
    ver2Arr := strings.Split(version2, ".")
    
    length := len(ver1Arr)
    if length < len(ver2Arr) {
        length = len(ver2Arr)
    }
    
    for i:=0; i < length; i++ {
        v1 := 0
        if i < len(ver1Arr) {
            v1, _ = strconv.Atoi(ver1Arr[i])
        }
        v2 := 0
        if i < len(ver2Arr) {
            v2, _ = strconv.Atoi(ver2Arr[i])
        }
        cp := 0
        if v1 > v2 {
            cp = 1
        } else if v2 > v1 {
            cp = -1
        }
        if cp != 0 {
            return cp
        }
    }
    return 0
}
```

## 实现

```golang
func compareVersion(version1 string, version2 string) int {
	var num1, num2 int
	var i, j int
	len1, len2 := len(version1), len(version2)
	for {
		// get each part of version1
		for ; i < len1; i++ {
			if version1[i] == '.' {
				break
			}
			num1 += 10*num1 + int(version1[i]-'0')
		}
		i++

		// get each part of version2
		for ; j < len2; j++ {
			if version2[j] == '.' {
				break
			}
			// cal sum
			num2 += 10*num2 + int(version2[j]-'0')
		}
		j++

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		} else {
			if j >= len2 && i >= len1 {
				break
			}
			num1, num2 = 0, 0
		}
	}
	return 0
}
```

## 改进
```golang

```

## 反思

## 参考