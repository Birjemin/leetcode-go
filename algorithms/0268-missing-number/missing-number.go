package missing_number

func missingNumber(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func missingNumber1(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	length := len(nums)
	return length * (length + 1) / 2 - sum
}

func missingNumber2(nums []int) int {
	var ret int
	for i, v := range nums {
		ret += i + 1 - v
	}
	return ret
}
