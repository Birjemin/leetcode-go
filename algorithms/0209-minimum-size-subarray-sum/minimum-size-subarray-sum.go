package minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	var sum int
	var count int
	var slow int

	for fast, v := range nums {
		sum += v
		count++

		// find !
		if sum >= s {
			for slow < fast && fast < length && count != 1 {
				sum -= nums[slow]
				slow++
				if sum >= s {
					count--
				} else if fast < length-1 {
					sum += nums[fast+1]
					fast++
				}
			}
			return count
		}
	}
	return 0
}

func minSubArrayLen1(s int, nums []int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	var (
		count int
		slow  int
		fast  int
		flag  bool
	)

	for i, v := range nums {
		s -= v

		// find !
		if s <= 0 {
			fast, count, flag = i, i+1, true
			break
		}
	}

	// not found
	if !flag {
		return 0
	}

	for slow < fast && fast < length && count != 1 {
		s += nums[slow]
		slow++
		if s <= 0 {
			count--
		} else if fast < length-1 {
			s -= nums[fast+1]
			fast++
		}
	}
	return count
}