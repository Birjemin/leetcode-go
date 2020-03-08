package restore_ip_addresses

import (
    "strings"
)

func restoreIpAddresses(s string) []string {
    var ret []string
    length := len(s)
    for i := 1; i < min(4, length); i++ {
        if !isValid(s[:i]) {
            continue
        }
        for j := i + 1; j < min(i+4, length); j++ {

            if !isValid(s[i:j]) {
                continue
            }
            for k := j + 1; k < min(j+4, length); k++ {
                if len(s[k:]) > 3 || !isValid(s[j:k]) || !isValid(s[k:]) {
                    continue
                }
                ret = append(ret, ip([]string{s[:i], s[i:j], s[j:k], s[k:]}))
            }
        }
    }
    if ret == nil {
        return []string{}
    }
    return ret
}

// it is a valid address
func isValid(s string) bool {

    switch len(s) {
    case 2:
        if s[0] == '0' {
            return false
        }
    case 3:
        if s[0] == '0' || s[0] > '2' {
            return false
        } else if s[0] == '2' {
            if s[1] > '5' {
                return false
            } else if s[1] == '5' && s[2] > '5' {
                return false
            }
        }
    }
    return true
}

// get ip's address
func ip(s []string) string {
    return strings.Join(s, ".")
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
