package palindrome_number

func isPalindrome1(x int) bool {
    if x < 0 {
        return false
    } else if x < 10 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    var y int
    for temp := x; temp != 0; temp /= 10 {
        y = 10 * y + temp % 10
    }
    if y == x {
        return true
    } else {
        return false
    }
}

func isPalindrome2(x int) bool {
    if x < 0 {
        return false
    } else if x < 10 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    for y := 0; x > y; {
        y = 10 * y + x % 10
        x /= 10
        // 只需要计算一半
        if x == y || x / 10 == y {
            return true
        }
    }
    return false
}