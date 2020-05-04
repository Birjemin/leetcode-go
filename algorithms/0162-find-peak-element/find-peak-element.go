package find_peak_element

func findPeakElement(nums []int) int {
    var idx int
    for i, v := range nums {
        if nums[idx] > v {
            break
        }
        idx = i
    }
    return idx
}

func findPeakElement1(nums []int) int {
    low, high := 0, len(nums)-1

    for low < high {
        mid := (low + high) / 2
        if nums[mid] < nums[mid+1] {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return high
}
