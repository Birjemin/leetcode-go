package length_of_last_word

import (
    "strings"
)

func lengthOfLastWord1(s string) int {
    var flag int
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == 32 {
            if flag > 0 {
                break
            }
        } else {
            flag++
        }
    }
    return flag
}

func lengthOfLastWord(s string) int {
    var count int
    s = strings.TrimSpace(s)
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == 32 {
            break
        }
        count++
    }
    return count
}
