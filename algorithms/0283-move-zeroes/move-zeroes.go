package move_zeroes

func moveZeroes(nums []int) {
	length, count := len(nums), 0
	for i := length - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}
		count++
		for j := i; j < length-count; j++ {
			nums[j] = nums[j+1]
		}
		nums[length-count] = 0
	}
}

func moveZeroes1(nums []int) {
	length := len(nums)
	var slow, fast int

	for fast < length {

		for slow != length-1 && nums[slow] != 0 {
			slow++
		}

		if fast < slow {
			fast = slow
		}

		for fast != length-1 && nums[fast] == 0 {
			fast++
		}

		nums[fast], nums[slow] = nums[slow], nums[fast]

		slow++
		fast++
	}
}

func moveZeroes2(nums []int) {
	var l int
	for i, v := range nums {
		if v == 0 {
			continue
		} else if i != l {
			nums[l], nums[i] = v, 0
		}
		l++
	}
}
