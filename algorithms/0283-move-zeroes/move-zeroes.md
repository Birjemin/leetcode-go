## 问题
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:
```
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:
```

- You must do this in-place without making a copy of the array.
- Minimize the total number of operations.

## 分析
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

常规方法移动，快慢指针，27、28、80题类似，这就很有趣~~
## 最高分
刘辟
```golang
func moveZeroes(nums []int)  {
	currentNotNullIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[currentNotNullIndex] = nums[i]

			if currentNotNullIndex != i {
				nums[i] = 0
			}

			currentNotNullIndex++
		}
	}
}
```

## 实现
最基本的实现
```golang
func moveZeroes(nums []int) {
	length, count := len(nums), 0
	for i := length - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}
		count++
		for j := i; j < length-count; j++ {
			nums[j] = nums[j+1]
		}
		nums[length-count] = 0
	}
}
```

## 改进
在循环中不断逼近，不断调节fast和slow
```golang
func moveZeroes(nums []int) {
	length := len(nums)
	var slow, fast int

	for fast < length {

		for slow != length-1 && nums[slow] != 0 {
			slow++
		}

		if fast < slow {
			fast = slow
		}

		for fast != length-1 && nums[fast] == 0 {
			fast++
		}

		nums[fast], nums[slow] = nums[slow], nums[fast]

		slow++
		fast++
	}
}
```

## 改进
使用27、80题
```golang
func moveZeroes(nums []int) {
	var l int
	for i, v := range nums {
		if v == 0 {
			continue
		} else if i != l {
			nums[l], nums[i] = v, 0
		}
		l++
	}
}

```

## 反思

## 参考