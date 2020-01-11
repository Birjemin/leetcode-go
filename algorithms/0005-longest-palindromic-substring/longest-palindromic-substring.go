package longest_palindromic_substring

func longestPalindrome1(s string) string {
    var res string
    var temp string
    length := len(s)
    var l, l2 int
    if length <= 1 {
        return s
    }
    for i := 0; i < length-1; i++ {
        for j := i + l; j <= length; j++ {
            temp = s[i:j]
            l2 = len(temp)
            if isPalindrome(temp, l2) {
                res = temp
                l = l2
            }
        }
    }
    return res
}

// 判断某一个是否是回文
func isPalindrome(s string, j int) bool {
    i := 0
    for i < j {
        if s[i] != s[j-1] {
            return false
        }
        i++
        j--
    }
    return true
}

func longestPalindrome(s string) string {
    min, max, length, l, h := 0, 0, len(s), 0, 0
    for i := 0; i < length; i++ {
        // 奇数
        l, h = i, i
        for l >= 0 && h < length && s[l] == s[h] {
            l--
            h++
        }
        if h-l-1 > max-min {
            min, max = l+1, h
        }
        // 偶数
        l, h = i, i+1
        for l >= 0 && h < length && s[l] == s[h] {
            l--
            h++
        }
        if h-l-1 > max-min {
            min, max = l+1, h
        }
    }
    return s[min:max]
}
