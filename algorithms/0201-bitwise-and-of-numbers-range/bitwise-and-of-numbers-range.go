package bitwise_and_of_numbers_range

func rangeBitwiseAnd(m int, n int) int {
	ret := m
	for i := m + 1; i <= n; i++ {
		ret &= i
	}
	return ret
}

func rangeBitwiseAnd1(m int, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	var count int

	for i := 31; i > 0; i-- {
		// if high bit is 1, find it!
		if m>>uint(i)&1 == 1 {
			count = i
		}

		if n>>uint(i)&1 == 1 {
			// because n >= m, count <= i
			// if count < i, return 0
			if count < i {
				return 0
			}
			break
		}
	}

	var ret int
	for i := count; i >= 0; i-- {

		// get high bit
		a1 := n >> uint(i)
		if a1 != m>>uint(i) {
			break
		}
		tmp := a1 << uint(i)
		ret |= tmp
		// subtract the high bit
		n -= tmp
		m -= tmp
	}
	return ret
}

func rangeBitwiseAnd2(m int, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	var ret int
	for i := 31; i >= 0; i-- {
		x := m >> uint(i)
		if x != n>>uint(i) {
			break
		}
		tmp := x << uint(i)
		ret |= tmp
		// subtract the high bit
		n -= tmp
		m -= tmp
	}
	return ret
}

func rangeBitwiseAnd3(m int, n int) int {

	var ret uint
	for n != m {
		n >>= 1
		m >>= 1
		ret++
	}
	return n << ret
}
