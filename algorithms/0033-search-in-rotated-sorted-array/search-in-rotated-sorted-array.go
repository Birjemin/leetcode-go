package search_in_rotated_sorted_array

func search(nums []int, target int) int {
    var idx int
    length := len(nums)
    // if array is null
    if length == 0 {
        return -1
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

func dichotomy(nums *[]int, start int, end int, target int) int {
    // unbelievable!!! end < start
    if end < start {
        return -1
    }
    div := (end - start) / 2
    if div == 0 {
        if target == (*nums)[end] {
            return end
        } else if target == (*nums)[start] {
            return start
        } else {
            return -1
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

func search1(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n - 1

    for l <= r {
        if nums[l] == target {
            return l
        }
        if nums[r] == target {
            return r
        }
        mid := (l + r) / 2
        if nums[mid] == target {
            return mid
        }
        // move l and r
        if nums[l] > nums[r] {
            l += 1
            r -= 1
        } else {
            if nums[mid] > target {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}
