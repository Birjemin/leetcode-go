package reverse_words_in_a_string

import (
    "strings"
)

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    s += " "

    var start int
    var ret string
    var flag bool
    for i, v := range s {
        if v != ' ' {
            flag = false
            continue
        }
        if !flag {
            ret = s[start:i] + " " + ret
            flag = true
        }
        start = i + 1
    }
    ret = strings.TrimSpace(ret)
    return ret
}

func reverseWords1(s string) string {

    // head and end space
    s = strings.TrimSpace(s)

    var tmp []string

    var start int

    // get array of vocabulary
    s += " "
    var flag bool
    var length int
    for i, v := range s {
        if v != ' ' {
            flag = false
            continue
        }
        if !flag {
            length++
            tmp = append(tmp, s[start:i])
            flag = true
        }
        start = i + 1
    }

    if len(tmp) > 0 {
        for i := 0; i < length/2; i++ {
            tmp[i], tmp[length-i-1] = tmp[length-i-1], tmp[i]
        }
    }

    return strings.Join(tmp, " ")
}

func reverseWords2(s string) string {
    var ret strings.Builder
    var flag bool
    s = " " + strings.TrimSpace(s)
    end := len(s)
    for i := end - 1; i >= 0; i-- {

        if s[i] != ' ' {
            flag = false
            continue
        }
        if !flag {
            ret.WriteString(s[i:end])
            flag = true
        }
        end = i
    }
    return strings.TrimSpace(ret.String())
}
