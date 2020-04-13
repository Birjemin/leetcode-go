## 问题
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
```
Input: [2,2,3,2]
Output: 3
```

Example 2:
```
Input: [0,1,0,1,0,1,99]
Output: 99
```

## 分析
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

136题可以使用位运算，那么这一题呢

## 最高分
想不到[url](https://cloud.tencent.com/developer/article/1131945)
```golang
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
```

## 实现
最简单的方式
```golang
func singleNumber(nums []int) int {
    ret := make(map[int]int, (len(nums)+2)/3)
    for _, v := range nums {
        if ret[v] >= 1 {
            ret[v]++
        } else {
            ret[v] = 1
        }
    }
    for k, v := range ret {
        if v == 1 {
            return k
        }
    }
    return 0
}
```

## 改进
使用一个32位数字，每一位代表一个数字，然后计数+1,最后为1的就是所求的值
```golang

```

## 反思

## 参考