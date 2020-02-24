## 问题
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
```
"123"
"132"
"213"
"231"
"312"
"321"
```
Given n and k, return the kth permutation sequence.

Note:

- Given n will be between 1 and 9 inclusive.
- Given k will be between 1 and n! inclusive.

Example 1:
```
Input: n = 3, k = 3
Output: "213"
```

Example 2:
```
Input: n = 4, k = 9
Output: "2314"
```

## 分析
f(1) = 1； n>1时，f(n) = n * f(n-1)。找到第k个元素，并且输出

## 最高分
```golang
func getPermutation(n int, k int) string {
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
```

## 实现
```golang
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

func reverse(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}

```

## 改进
最高分的解法很棒，先把k--，这样就不需要再判断整除和余数的问题了，不然的话需要将字符串翻转
```golang
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
```

## 反思

## 参考