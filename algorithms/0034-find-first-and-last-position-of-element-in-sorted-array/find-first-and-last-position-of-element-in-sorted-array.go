package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    length := len(nums)
    // if nums is invalid
    if length == 0 {
        return res
    }
    // find first position of nums
    searchLeft(&nums, 0, length-1, target, &res[0])
    // could not find the first position of nums
    if res[0] == -1 {
        return res
    }
    // find last position of nums
    searchRight(&nums, res[0], length-1, target, &res[1])
    return res
}

func searchLeft(nums *[]int, start int, end int, target int, res *int) {
    // find it!
    if target == (*nums)[start] {
        *res = start
        return
    }
    div := (end - start) / 2
    // that's mean the index of end is target value or could not found the index
    if div <= 0 {
        if target == (*nums)[end] {
            *res = end
        }
        return
    }
    if target <= (*nums)[div+start] {
        searchLeft(nums, start, end-div, target, res)
    } else {
        searchLeft(nums, start+div, end, target, res)
    }
}

func searchRight(nums *[]int, start int, end int, target int, res *int) {
    if target == (*nums)[end] {
        *res = end
        return
    }
    div := (end - start) / 2
    if div <= 0 {
        if target == (*nums)[start] {
            *res = start
        }
        return
    }
    if target < (*nums)[div+start] {
        searchRight(nums, start, end-div, target, res)
    } else {
        searchRight(nums, start+div, end, target, res)
    }
}
