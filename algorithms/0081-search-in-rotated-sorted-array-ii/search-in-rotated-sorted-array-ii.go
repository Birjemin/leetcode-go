package search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
    var idx int
    length := len(nums)
    // if array is null
    if length == 0 {
        return false
    }
    // search the reverse position
    for ; idx < length-1; idx++ {
        if nums[idx] > nums[idx+1] {
            break
        }
    }
    // handle
    if target >= nums[0] {
        // search the front part of array
        return dichotomy(&nums, 0, idx, target)
    } else {
        // search the next part of array
        return dichotomy(&nums, idx+1, length-1, target)
    }
}

func dichotomy(nums *[]int, start int, end int, target int) bool {
    // unbelievable!!! end < start
    if end < start {
        return false
    }
    div := (end - start) / 2
    if div == 0 {
        if target == (*nums)[end] || target == (*nums)[start] {
            return true
        } else {
            return false
        }
    }
    // search the front part of array
    if (*nums)[start+div] > target {
        return dichotomy(nums, start, end-div, target)
    } else {
        // search the next part of array
        return dichotomy(nums, start+div, end, target)
    }
}