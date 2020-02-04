package zigzag_conversion

import "bytes"

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

func convert1(s string, numRows int) string {
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
