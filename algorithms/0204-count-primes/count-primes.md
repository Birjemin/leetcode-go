## 问题
Count the number of prime numbers less than a non-negative number, n.

Example:
```
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
```

## 分析
统计所有小于非负整数 n 的质数的数量。

没啥难度，只要记住质数的计算方法(在一般领域，对正整数n，如果用2到开根号2之间的所有整数去除，均无法整除，则n为质数。)

## 最高分
```golang
func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}
```

## 实现
速度略慢
```golang
func countPrimes(n int) int {
	var ret int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ret++
		}
	}
	return ret
}

func isPrime(n int) bool {
	t := int(math.Sqrt(float64(n)))
	for i := 2; i <= t; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```

## 改进
厄拉多塞筛法：要得到自然数n以内的全部质数，必须把不大于根号n的所有质数的倍数剔除，剩下的就是质数。
```golang
func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}
	var ret int
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			ret++
		}
	}

	return ret
}
```

## 反思

## 参考