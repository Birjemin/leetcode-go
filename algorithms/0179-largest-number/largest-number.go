package largest_number

import (
	"strconv"
	"strings"
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

// quick's sort
func largestNumber1(nums []int) string {
	strs := make([]string, len(nums))

	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	quickSort(strs, 0, len(strs)-1)

	ret := strings.Join(strs, "")
	for ret[0] == '0' && len(ret) > 1 {
		ret = ret[1:]
	}

	return ret
}

func quickSort(src []string, first, last int) {
	flag, left, right := first, first, last

	if first >= last {
		return
	}

	for first < last {
		// left part
		for first < last {
			// do not move
			if src[flag]+src[last] >= src[last]+src[flag] {
				last -= 1
				continue
			} else {
				// move
				src[last], src[flag] = src[flag], src[last]
				flag = last
				break
			}
		}
		// right part
		for first < last {
			if src[flag]+src[first] <= src[first]+src[flag] {
				first += 1
				continue
			} else {
				src[first], src[flag] = src[flag], src[first]
				flag = first
				break
			}
		}
	}

	quickSort(src, left, flag-1)
	quickSort(src, flag+1, right)
}