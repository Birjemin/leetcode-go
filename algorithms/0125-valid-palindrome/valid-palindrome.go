package valid_palindrome

func isPalindrome(s string) bool {
    for left, right := 0, len(s)-1; left < right; {
        // to lower
        l, r := toLower(s[left]), toLower(s[right])

        if isNotString(l) {
            left++
            continue
        }
        if isNotString(r) {
            right--
            continue
        }

        if l != r {
            return false
        }
        left++
        right--
    }

    return true
}

func isNotString(s byte) bool {
    if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
        return false
    }
    return true
}

func toLower(s byte) byte {
    if s >= 'A' && s <= 'Z' {
        return s + 32
    }
    return s
}
