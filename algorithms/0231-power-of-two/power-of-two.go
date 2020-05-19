package power_of_two

func isPowerOfTwo(n int) bool {
	a := 1
	for a < n {
		a <<= 1
	}
	if n == a {
		return true
	}
	return false
}

func isPowerOfTwo1(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n % 2 == 1 {
			return false
		}
		n >>= 1
	}
	return true
}

