## 问题
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
```
Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
```

Example 2:
```
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
```

## 分析
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

198的加强版本，这题是环形，198是线性

## 最高分
```golang
func rob(nums []int) int {
    
    calc := func (arr []int) float64 {
        var prevMax, max float64
        for _, v := range arr {
            prevMax, max = max, math.Max(max, prevMax + float64(v))
        }
        return max
    }
    
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0] 
    }
    
    var try0, try1 float64
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() { defer wg.Done(); try0 = calc(nums[:len(nums)-1]) }() // w/o tail
    go func() { defer wg.Done(); try1 = calc(nums[1:]) }() // w/o head
    wg.Wait()
    
    return int(math.Max(try0, try1))
}
```

## 实现
没想出来，试图在198上面想办法，但是没想出来，看了提示，知道得到[1,n]的最大值，[0, n-1]最大值，然后比较
```golang
func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	one, two := cal(nums, 1, length), cal(nums, 0, length-1)
	if one > two {
		return one
	}
	return two
}

func cal(nums []int, start, length int) int {
	currMax, prevMax := 0, 0
	for i := start; i < length; i++ {
		temp := currMax
		if prevMax+nums[i] > currMax {
			currMax = prevMax + nums[i]
		}
		prevMax = temp
	}
	return currMax
}
```

## 改进
使用golang的特性，协程
```golang
func cal(nums []int, start, length int) int {
	currMax, prevMax := 0, 0
	for i := start; i < length; i++ {
		temp := currMax
		if prevMax+nums[i] > currMax {
			currMax = prevMax + nums[i]
		}
		prevMax = temp
	}
	return currMax
}

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	var wg sync.WaitGroup
	one, two := 0, 0

	wg.Add(2)

	go func() {
		one = cal(nums, 1, length)
		wg.Done()
	}()

	go func() {
		two = cal(nums, 0, length-1)
		wg.Done()
	}()

	wg.Wait()

	if one > two {
		return one
	}
	return two
}
```

## 反思

## 参考