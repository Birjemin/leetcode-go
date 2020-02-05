## 问题
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

Example:
```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

## 分析
求最大面积问题，最简单的方式就是遍历，考察的是双指针法的使用

## 最高分
```golang

```


## 实现
最直接的方式，也是性能很差的方式
```golang
func maxArea(height []int) int {
    var max, sum int
    length := len(height)
    for i := 0; i < length; i++ {
        for j := length - 1; j > i; j-- {
            sum = min(height[i], height[j]) * (j - i)
            if sum > max {
                max = sum
            }
        }
    }
    return max
}

func min(n1, n2 int) int {
    if n1 > n2 {
        return n2
    } else {
        return n1
    }
}
```

## 改进
使用双指针法：影响面积的因素是宽和高，如果左边短，则右移找更高，如果右边短，则左移找更高（宽度减少）
```golang
func maxArea(height []int) int {
    i, j := 0, len(height)-1
    if j == 0 {
        return 0
    }
    var max, area, h int
    // check
    for i < j {
        // left
        if height[i] < height[j] {
            h = height[i]
            i++
        } else {
            // right
            h = height[j]
            j--
        }
        //
        area = (j - i + 1) * h
        if area > max {
            max = area
        }
    }
    return max
}
```

## 反思

## 参考
