package sort_colors

func sortColors(nums []int) {
    right := len(nums) - 1
    if right <= 0 {
        return
    }
    var left, curr int
    for curr <= right {
        switch nums[curr] {
        case 0:
            nums[left], nums[curr] = nums[curr], nums[left]
            left++
            curr++
        case 2:
            nums[curr], nums[right] = nums[right], nums[curr]
            right--
        default:
            curr++
        }
    }
    return
}
