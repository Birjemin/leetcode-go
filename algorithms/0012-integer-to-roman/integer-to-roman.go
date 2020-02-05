package integer_to_roman

import (
    "bytes"
)

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