package longest_common_prefix

func longestCommonPrefix1(strs []string) string {
    length := len(strs)
    if length == 0 {
        return ""
    } else if length == 1 {
        return strs[0]
    }

    var flag string
    for i, s := range strs[0] {
        for j := 1; j < length; j++ {
            if len(strs[j]) <= i || strs[j][i] != byte(s) {
                return flag
            }
        }
        flag += string(s)
    }
    return flag
}

func longestCommonPrefix2(strs []string) string {
    length := len(strs)
    if length == 0 {
        return ""
    } else if length == 1 {
        return strs[0]
    }

    var flag []byte
    lenMap := make(map[int]int, length - 1)
    for i, s := range strs[0] {
        for j := 1; j < length; j++ {
            if lenMap[j] == 0 {
                lenMap[j] = len(strs[j])
            }
            if lenMap[j] <= i || strs[j][i] != byte(s) {
                return string(flag)
            }
        }
        flag = append(flag, byte(s))
    }
    return string(flag)
}