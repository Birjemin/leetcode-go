package count_and_say

import (
    "strconv"
    "strings"
)

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    } else {
        str := countAndSay(n - 1)
        var flag rune
        var count int
        var res strings.Builder
        for _, s := range str {
            if flag == 0 {
                flag = s
            }
            if flag == s {
                count++
            } else {
                res.WriteString(strconv.Itoa(count))
                res.WriteRune(flag)
                flag = s
                count = 1
            }
        }
        res.WriteString(strconv.Itoa(count))
        res.WriteRune(flag)
        return res.String()
    }
}
