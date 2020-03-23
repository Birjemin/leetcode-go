package rotate_array

func rotate(nums []int, k int) {
    length := len(nums)
    if length == 1 {
        return
    }
    for i := 0; i < k; i++ {
        for j := length; j > 1; j-- {
            nums[j-1], nums[j-2] = nums[j-2], nums[j-1];
        }
    }
}

func rotate1(nums []int, k int) {
    length := len(nums)
    if length < 2 {
        return
    }
    if k > length {
        k = k % length
    }
    reverse(nums, length)
    reverse(nums[:k], k)
    reverse(nums[k:], length-k)
}

func reverse(nums []int, length int) {
    for i := 0; i < length/2; i++ {
        nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
    }
}
