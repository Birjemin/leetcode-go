## 问题
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

Example:
```
Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
```

## 分析
最简单的方式还是直接循环遍历
- 使用dfs搜索

## 最高分
```golang
func restoreIPAddresses(s string) []string {
    if s == "" {
        return []string{}
    }
    res, ip := []string{}, []int{}
    dfs(s, 0, ip, &res)
    return res
}

func dfs(s string, index int, ip []int, res *[]string) {
    if index == len(s) {
        if len(ip) == 4 {
            *res = append(*res, getString(ip))
        }
        return
    }
    if index == 0 {
        num, _ := strconv.Atoi(string(s[0]))
        ip = append(ip, num)
        dfs(s, index+1, ip, res)
    } else {
        num, _ := strconv.Atoi(string(s[index]))
        next := ip[len(ip)-1]*10 + num
        if next <= 255 && ip[len(ip)-1] != 0 {
            ip[len(ip)-1] = next
            dfs(s, index+1, ip, res)
            ip[len(ip)-1] /= 10
        }
        if len(ip) < 4 {
            ip = append(ip, num)
            dfs(s, index+1, ip, res)
            ip = ip[:len(ip)-1]
        }
    }
}

func getString(ip []int) string {
    res := strconv.Itoa(ip[0])
    for i := 1; i < len(ip); i++ {
        res += "." + strconv.Itoa(ip[i])
    }
    return res
}
```

## 实现
最直接的方式实现，暴力破解
```golang
func restoreIpAddresses(s string) []string {
    var ret []string
    length := len(s)
    for i := 1; i < min(4, length); i++ {
        if !isValid(s[:i]) {
            continue
        }
        for j := i + 1; j < min(i+4, length); j++ {
            if !isValid(s[i:j]) {
                continue
            }
            for k := j + 1; k < min(j+4, length); k++ {
                if len(s[k:]) > 3 || !isValid(s[j:k]) || !isValid(s[k:]) {
                    continue
                }
                ret = append(ret, ip([]string{s[:i], s[i:j], s[j:k], s[k:]}))
            }
        }
    }
    if ret == nil {
        return []string{}
    }
    return ret
}

// it is a valid address
func isValid(s string) bool {
    switch len(s) {
    case 2:
        if s[0] == '0' {
            return false
        }
    case 3:
        if s[0] == '0' || s[0] > '2' {
            return false
        } else if s[0] == '2' {
            if s[1] > '5' {
                return false
            } else if s[1] == '5' && s[2] > '5' {
                return false
            }
        }
    }
    return true
}

// get ip's address
func ip(s []string) string {
    return strings.Join(s, ".")
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

```

## 改进
可以使用dfs方式遍历，但是有一个注意事项就是：出口条件，4层，感觉还是for处理吧。
```golang

```

## 反思

## 参考