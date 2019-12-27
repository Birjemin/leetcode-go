package roman_to_integer

func romanToInt1(s string) int {
    table := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    length := len(s)
    sum := table[rune(s[length - 1])]
    for i := length - 1; i > 0; i-- {
        z := table[rune(s[i-1])]
        if table[rune(s[i])] > z {
            sum -= z
        } else {
            sum += z
        }
    }
    return sum
}

func romanToInt2(s string) int {
    table := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    if s == "" {
        return 0
    }
    var sum, next int
    for i := len(s); i > 0; i-- {
        temp := table[rune(s[i-1])]
        if next > temp {
            sum -= temp
        } else {
            sum += temp
        }
        next = temp
    }
    return sum
}