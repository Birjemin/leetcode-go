package excel_sheet_column_number

func titleToNumber(s string) int {
    rate, sum := 1, 0
    for i := len(s) - 1; i >= 0; i-- {
        sum += rate * (int(s[i]-'A'+1))
        rate *= 26
    }
    return sum
}

func titleToNumber1(s string) int {
    var sum int
    for _, v := range s {
        sum = 26*sum + int(v-'A'+1)
    }
    return sum
}
