package plus_one

func plusOne(digits []int) []int {
    // 处理进位
    for i := len(digits) - 1; i >= 0; i-- {
        // 当前位不会产生进位
        if digits[i] < 9 {
            // 如果i = len(digits) - 1则代表+1，反之代表进位1
            digits[i] += 1
            return digits
        }
        // 如果digits[i] == 9 则产生进位1，当前位为0
        digits[i] = 0
    }
    // 处理首部进位
    return append([]int{1}, digits...)
}
