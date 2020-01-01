package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
    left, length := 0, len(nums)
    if length == 0 {
        return 0
    }
    for right := 1; right < length; right++ {
        if nums[left] < nums[right] {
            left++
            nums[left], nums[right] = nums[right], nums[left]
        }
    }
    return left + 1
}