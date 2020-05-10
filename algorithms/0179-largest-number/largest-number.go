package largest_number

import (
	"strconv"
)

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
