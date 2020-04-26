package evaluate_reverse_polish_notation

import (
    "strconv"
)

func evalRPN(tokens []string) int {
    var tmp []int
    var length int
    for _, v := range tokens {
        // check the string is number
        if num, err := strconv.Atoi(v); err == nil {
            tmp = append(tmp, num)
            length++
        } else {
            operate(&tmp, &length, v)
            length--
        }
    }
    return tmp[0]
}

func operate(ret *[]int, length *int, operate string) {
    num := (*ret)[*length-1]
    *ret = (*ret)[:*length-1]
    switch operate {
    case "+":
        (*ret)[*length-2] += num
    case "-":
        (*ret)[*length-2] -= num
    case "*":
        (*ret)[*length-2] *= num
    case "/":
        (*ret)[*length-2] /= num
    }
}
