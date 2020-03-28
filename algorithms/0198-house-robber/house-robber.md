## 问题
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
```
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
```

Example 2:
```
Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
```

## 分析
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

可以推导出公式 f(n) = {f(n-1), f(n-2)+n}max

使用dp算法吧

## 最高分
```golang
func rob(nums []int) int {
    prevMax := 0
    currMax := 0
    
    for i:=0; i < len(nums); i++ {
        temp:=currMax
        if prevMax + nums[i] > currMax {
            currMax = prevMax + nums[i] 
        }
        prevMax = temp
    }
    return currMax
}
```

## 实现
```golang
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

```

## 改进
在原来的地址上改动，节省空间(和dp算法类似)
```golang
func rob(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    } else if length == 1 {
        return nums[0]
    }
    cal(&nums, length-1)
    return nums[length-1]
}

func cal(dp *[]int, i int) {
    if i == 1 {
        if (*dp)[i] < (*dp)[0] {
            (*dp)[i] = (*dp)[0]
        }
        return
    }

    cal(dp, i-1)

    t := (*dp)[i-2] + (*dp)[i]
    if t > (*dp)[i-1] {
        (*dp)[i] = t
    } else {
        (*dp)[i] = (*dp)[i-1]
    }
}

```

## 反思

## 参考