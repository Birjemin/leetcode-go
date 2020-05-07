## 问题
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

Example 1:
```
Input: numerator = 1, denominator = 2
Output: "0.5"
```

Example 2:
```
Input: numerator = 2, denominator = 1
Output: "2"
```

Example 3:
```
Input: numerator = 2, denominator = 3
Output: "0.(6)"
```

## 分析
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
如果小数部分为循环小数，则将循环的部分括在括号内。
- 如果判断循环小数？？
- 如何使用buffer？？

## 最高分
```golang
func fractionToDecimal(numerator int, denominator int) string {
	flag := ""
	switch {
	case numerator < 0 && denominator < 0:
		numerator, denominator = -numerator, -denominator
	case numerator > 0 && denominator < 0:
		denominator, flag = -denominator, "-"
	case numerator < 0 && denominator > 0:
		numerator, flag = -numerator, "-"
	}

	qua, rem := numerator/denominator, numerator%denominator
	res := []byte(strconv.FormatInt(int64(qua), 10))
	if rem == 0 {
		return flag + string(res)
	}

	res = append(res, '.')
	m, curIndex := make(map[int]int), len(res)-1
	for rem != 0 {
		numerator = rem * 10
		qua, rem = numerator/denominator, numerator%denominator

		if index, ok := m[numerator]; ok {
			res = append(res, ')', ' ')
			copy(res[index+1:], res[index:])
			res[index] = '('
			break
		}

		res = append(res, digitToByte(qua))
		curIndex++
		m[numerator] = curIndex
	}
	return flag + string(res)
}

func digitToByte(digit int) byte {
	return byte('0' + digit)
}
```

## 实现
不停的查找边界，心酸
```golang
func fractionToDecimal(numerator int, denominator int) string {
	var tmp []string
	kv := make(map[int]int)
	var count int
	var flag bool
	var negative bool
	var ret string

	if numerator < 0 {
		negative = true
		numerator = -numerator
	}

	if denominator < 0 {
		if negative {
			negative = false
		} else {
			negative = true
		}
		denominator = -denominator
	}

	for {

		divider, remainder := numerator/denominator, numerator%denominator
		tmp = append(tmp, strconv.Itoa(divider))
		if remainder == 0 {
			ret = strings.Join(tmp, "")
			break
		}


		// decimal fraction
		if !flag && remainder < denominator {
			count++
			flag = true
			tmp = append(tmp, ".")
		}

		if v, ok := kv[remainder]; ok {
			ret = strings.Join(tmp[:v+1], "") + "(" + strings.Join(tmp[v+1:], "") + ")"
			break
		}

		kv[remainder] = count

		numerator = 10 * remainder
		count++
	}

	if ret != "0" && negative {
		return "-" + ret
	}
	return ret
}

```

## 改进
```golang
func fractionToDecimal(numerator int, denominator int) string {
	var tmp []string
	var count int
	var flag, negative bool
	var ret string
	kv := make(map[int]int)

	// check if the number is negative
	negative = checkNegative(&numerator, &denominator)

	for {

		divider, remainder := numerator/denominator, numerator%denominator

		tmp = append(tmp, strconv.Itoa(divider))

		if remainder == 0 {
			ret = strings.Join(tmp, "")
			break
		}

		// decimal fraction
		if !flag && remainder < denominator {
			count++
			flag = true
			tmp = append(tmp, ".")
		}

		if v, ok := kv[remainder]; ok {
			ret = strings.Join(tmp[:v+1], "") + "(" + strings.Join(tmp[v+1:], "") + ")"
			break
		}

		kv[remainder] = count

		numerator = 10 * remainder
		count++
	}

	if ret != "0" && negative {
		return "-" + ret
	}
	return ret
}

// check if the number is negative
func checkNegative(numerator *int, denominator *int) (negative bool) {
	if *numerator < 0 {
		negative = true
		*numerator = -*numerator
	}

	if *denominator < 0 {
		if negative {
			negative = false
		} else {
			negative = true
		}
		*denominator = -*denominator
	}
	return
}

```

## 反思

## 参考