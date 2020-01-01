package implement_strstr

import "strings"

func strStr1(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}

func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    length, l := len(haystack), len(needle)
    for idx := 0; idx <= length-l; idx++ {
        if haystack[idx:l+idx] == needle {
            return idx
        }
    }
    return -1
}
