package permutation_sequence

import (
    "strings"
)

func getPermutation(n int, k int) string {
    if n == 1 {
        return "1"
    }
    var str string
    for i := 0; i < n; i++ {
        str += string(i + '1')
    }
    ret := cal(n, k, str, "")
    return ret
}

func cal(n int, k int, str string, ret string) string {
    // if
    sub := getNum(n)
    quotient, remainder := k/sub, k%sub
    // got it
    if remainder == 0 {
        ret += string(str[quotient-1])
        str = strings.Replace(str, string(str[quotient-1]), "", 1)
        return ret + reverse(str)
    }
    // find next value
    ret += string(str[quotient])
    str = strings.Replace(str, string(str[quotient]), "", 1)
    return cal(n-1, remainder, str, ret)
}

func getNum(n int) int {
    if n == 1 {
        return 1
    }
    return (n - 1) * getNum(n-1)
}

// reverse string
func reverse(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}

func getPermutation1(n int, k int) string {
    if n == 0 {
        return ""
    }
    data := make([]int, n)
    data[0] = 1
    iToS := []byte("123456789")
    for i := 1; i < n; i++ {
        data[i] = data[i-1] * i
    }
    ret := ""
    k--
    for i := n - 1; i >= 0; i-- {
        ret += string(iToS[k/data[i]])
        iToS = append(iToS[:k/data[i]], iToS[k/data[i]+1:]...)
        k = k % data[i]
    }
    return ret
}
