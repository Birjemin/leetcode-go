## 问题
Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:
```
Input: [10,2]
Output: "210"
```

Example 2:
```
Input: [3,30,34,5,9]
Output: "9534330"
```

Note: The result may be very large, so you need to return a string instead of an integer.

## 分析
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
其实就是比较每个数字的最高位，最高的排前面就行了，使用buffer保存结果

想了一天，想不出来，看似很简单，但是解法还真的难想~(golang中有字符串的比较大小)

## 最高分
```golang
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(ByLargestNumber(strs))

	var res bytes.Buffer
	for i := 0; i < len(strs); i++ {
		res.WriteString(strs[i])
	}
	strRes := res.String()
	if strRes[0] == '0' {
		return "0"
	}
	return strRes
}

type ByLargestNumber []string

func (n ByLargestNumber) Len() int {
	return len(n)
}
func (n ByLargestNumber) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n ByLargestNumber) Less(i, j int) bool {
	return !compare(n[i] + n[j], n[j] + n[i])
}

func compare(a, b string) bool {
	for i := 0; i < len(a); i++ {
		switch {
		case a[i] < b[i]:
			return true
		case a[i] > b[i]:
			return false
		}
	}
	return true
}
```

## 实现
需要找出一个规律，这个规律就是插入一个数字a到str中，如果`a+str > str[:idx]+a+str[idx:]`，则说明需要插入到该位置！
```golang
func largestNumber(nums []int) string {
	var ret string
	var queue []int

	for _, num := range nums {
		t := strconv.Itoa(num)
		flag, idx := false, 0

		for i, q := range queue {
			if t+ret[idx:] > ret[idx:idx+q]+t+ret[idx+q:] {
				ret = ret[:idx] + t + ret[idx:]
				queue = append(queue[:i], append([]int{len(t)}, queue[i:]...)...)
				flag = true
				break
			}
			idx += q
		}

		if !flag {
			queue = append(queue, len(t))
			ret += t
		}
	}

	for ret[0] == '0' && len(ret) > 1 {
		ret = ret[1:]
	}

	return ret
}

```

## 改进
```golang
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(ByLargestNumber(strs))

	var res bytes.Buffer
	for _, str := range strs {
		res.WriteString(str)
	}
	strRes := res.String()

	if strRes[0] == '0' {
		return "0"
	}
	return strRes
}

type ByLargestNumber []string

func (n ByLargestNumber) Len() int {
	return len(n)
}
func (n ByLargestNumber) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n ByLargestNumber) Less(i, j int) bool {
	return n[i]+n[j] > n[j]+n[i]
}

```

## 反思

## 参考