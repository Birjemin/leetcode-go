package letter_combinations_of_a_phone_number

var (
    m = [][]string{
        {"a", "b", "c"},
        {"d", "e", "f"},
        {"g", "h", "i"},
        {"j", "k", "l"},
        {"m", "n", "o"},
        {"p", "q", "r", "s"},
        {"t", "u", "v"},
        {"w", "x", "y", "z"},
    }
    dic = []string{
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "wxyz",
    }
)

func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    return find(digits, []string{})
}

func find(str string, s []string) []string {
    var ret []string
    if str == "" {
        return s
    }
    length := len(s)
    for _, v := range m[str[0]-'0'-2] {
        if length == 0 {
            ret = append(ret, string(v))
        } else {
            for _, o := range s {
                ret = append(ret, string(o)+string(v))
            }
        }
    }
    return find(str[1:], ret)
}



