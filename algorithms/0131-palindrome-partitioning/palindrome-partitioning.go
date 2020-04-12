package palindrome_partitioning

func partition(s string) [][]string {
    var ret [][]string
    dfs([]string{}, s, &ret)
    return ret
}

func dfs(prefix []string, s string, ret *[][]string) {
    if s == "" {
        b := make([]string, len(prefix))
        copy(b, prefix)
        *ret = append(*ret, b)
        return
    }

    for i := range s {
        if isPalindrome(i, s) {
            dfs(append(prefix, s[:i+1]), s[i+1:], ret)
        }
    }
}

// palindrome
func isPalindrome(end int, s string) bool {
    for i := 0; i < end; i, end = i+1, end-1 {
        if s[i] != s[end] {
            return false
        }
    }
    return true
}
