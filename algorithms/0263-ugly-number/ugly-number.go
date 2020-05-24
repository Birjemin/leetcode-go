package ugly_number

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for {
		if num == 1 {
			return true
		}
		a1, a2, a3 := num%2, num%3, num%5
		if a1 != 0 && a2 != 0 && a3 != 0 {
			return false
		}
		if a1 == 0 {
			num /= 2
		}
		if a2 == 0 {
			num /= 3
		}
		if a3 == 0 {
			num /= 5
		}
	}
}

func isUgly1(num int) bool {
	if num <= 0 {
		return false
	}

	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}

func isUgly2(num int) bool {
	if num <= 0 {
		return false
	}

	if num%2 == 0 {
		return isUgly2(num / 2)
	}
	if num%3 == 0 {
		return isUgly2(num / 3)
	}
	if num%5 == 0 {
		return isUgly2(num / 5)
	}
	return num == 1
}
