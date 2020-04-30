package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
    for i, v := range nums {
        if i == 0 {
            continue
        }
        if v < nums[i-1] {
            return v
        }
    }
    return nums[0]
}

func findMin1(nums []int) int {
    left, right := 0, len(nums)-1
    for {

        if right-left == 0 {
            return nums[0]
        } else if right-left == 1 {
            if nums[left] < nums[right] {
                return nums[0]
            }
            return nums[right]
        }

        // split the array
        mid := (left + right) / 2
        if nums[left] > nums[mid] {
            right = mid
        } else {
            left = mid
        }
    }
}
