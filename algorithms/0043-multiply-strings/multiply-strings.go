package multiply_strings

import (
    "bytes"
)

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
