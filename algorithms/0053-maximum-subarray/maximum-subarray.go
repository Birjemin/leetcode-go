package maximum_subarray

func maxSubArray1(nums []int) int {
    length := len(nums)
    if length == 1 {
        return nums[0]
    }
    flag := nums[0]

    for idx, n := range nums {
        // 当前循环n大于上一次循环的flag，则flag更新
        if flag < n {
            flag = n
        }
        // 在当前循环中找到最大值flag
        for right := idx + 1; right < length; right++ {
            n += nums[right]
            if flag < n {
                flag = n
            }
        }
    }
    return flag
}

func maxSubArray(nums []int) int {
    max, sum := nums[0], nums[0]
    for _, v := range nums[1:] {
        if sum < 0 {
            sum = v
        } else {
            sum += v
        }
        if max < sum {
            max = sum
        }
    }
    return max
}
