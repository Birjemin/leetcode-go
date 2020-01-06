package add_binary

func addBinary1(a string, b string) string {
    if a == "" && b == "" {
        return "0"
    }
    l1, l2 := len(a), len(b)
    // 先保证a的长度大于b的长度
    if l1 < l2 {
        a, b = b, a
        l1, l2 = l2, l1
    }
    var y, sum int
    ret := make([]byte, l1+1)
    for i := 0; i < l1; i++ {
        if i < l2 {
            y += byte2Int(b[l2-i-1])
        }
        sum = byte2Int(a[l1-1-i]) + y
        if sum == 1 {
            y = 0
            ret[l1-i] = 49
        } else if sum == 2 {
            y = 1
            ret[l1-i] = 48
        } else if sum == 3 {
            y = 1
            ret[l1-i] = 49
        } else {
            y = 0
            ret[l1-i] = 48
        }
    }
    if y == 1 {
        ret[0] = 49
    } else {
        ret = ret[1:]
    }
    return string(ret)
}

func byte2Int(s byte) int {
    if s == 49 {
        return 1
    }
    return 0
}

func addBinary(a string, b string) string {
    if a == "" && b == "" {
        return "0"
    }
    l1, l2 := len(a), len(b)
    // 先保证a的长度大于b的长度
    if l1 < l2 {
        a, b = b, a
        l1, l2 = l2, l1
    }
    var flag, x, y int
    ret := make([]byte, l1+1)
    for i := 0; i < l1; i++ {
        if i < l2 {
            x = byte2Int(b[l2-i-1])
        } else {
            x = 0
        }
        y = byte2Int(a[l1-i-1])
        // 两个不相等，则会改变当前位
        ret[l1-i] = byte(x ^ y ^ flag + 48)
        // 两个相等，则会产生移位
        flag = x&y | flag&(x^y)
    }
    if flag == 1 {
        ret[0] = 49
    } else {
        ret = ret[1:]
    }
    return string(ret)
}
