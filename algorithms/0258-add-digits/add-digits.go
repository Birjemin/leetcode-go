package add_digits

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if ret := num % 9; ret == 0 {
		return 9
	} else {
		return ret
	}
}

func addDigits1(num int) int {
	return (num-1)%9 + 1
}
