package happy_number

func isHappy(n int) bool {
	var sum, remainder int
	for {
		for {
			remainder = n % 10
			sum += remainder * remainder
			n /= 10

			if n == 0 {
				n, sum = sum, 0
				break
			}
		}

		if n == 1 {
			return true
		} else if n == 4 {
			return false
		}
	}
}

func isHappy1(n int) bool {
	slow, fast := n, cal(n)
	for slow != fast {
		slow,fast = cal(slow),cal(cal(fast))
	}
	if slow == 1 {
		return true
	}
	return false
}

func cal(n int) int {
	var sum, remainder int
	for n!=0 {
		remainder = n % 10
		sum += remainder * remainder
		n /= 10
	}
	return sum
}