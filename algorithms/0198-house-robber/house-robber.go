package house_robber

func rob(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    } else if length == 1 {
        return nums[0]
    }
    dp := make([]int, length)
    cal(&dp, nums, length-1)
    return dp[length-1]
}

func cal(dp *[]int, nums []int, i int) {
    if i == 1 {
        (*dp)[0] = nums[0]
        if nums[i] < nums[0] {
            (*dp)[i] = nums[0]
        } else {
            (*dp)[i] = nums[i]
        }
        return
    }

    cal(dp, nums, i-1)

    t := (*dp)[i-2] + nums[i]
    if t > (*dp)[i-1] {
        (*dp)[i] = t
    } else {
        (*dp)[i] = (*dp)[i-1]
    }
}


func rob1(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    } else if length == 1 {
        return nums[0]
    }
    cal1(&nums, length-1)
    return nums[length-1]
}

func cal1(dp *[]int, i int) {
    if i == 1 {
        if (*dp)[i] < (*dp)[0] {
            (*dp)[i] = (*dp)[0]
        }
        return
    }

    cal1(dp, i-1)

    t := (*dp)[i-2] + (*dp)[i]
    if t > (*dp)[i-1] {
        (*dp)[i] = t
    } else {
        (*dp)[i] = (*dp)[i-1]
    }
}
