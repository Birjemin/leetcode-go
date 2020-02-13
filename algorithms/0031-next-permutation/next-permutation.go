package next_permutation

func nextPermutation(nums []int) {
    length := len(nums)
    p, j := length-1, 0
    // find position
    for ; p > 0; p-- {
        if nums[p-1] < nums[p] {
            break
        }
    }
    // reverse the rest of array
    for ; length-p > j*2; j++ {
        nums[p+j], nums[length-j-1] = nums[length-j-1], nums[p+j]
    }
    // all array had been reversed
    if p == 0 {
        return
    }

    // find position for the next array
    for j = 0; nums[p-1] >= nums[p+j]; j++ {
    }
    nums[p+j], nums[p-1] = nums[p-1], nums[p+j]
    return
}
