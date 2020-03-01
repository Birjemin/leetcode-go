package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 2 {
        return length
    }
    j := 1
    for i := 2; i < length; i++ {
        // equal
        if nums[j-1] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }
    return j + 1
}
