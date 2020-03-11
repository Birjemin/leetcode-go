## 问题
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:
```
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

## 分析
和95题目一样，找规律的题目C(n+1) = [2*(2*n+1)/(n+2)]*C(n)

## 最高分
其实都一样~~
```golang
func numTrees(n int) int {
    trees := make([]int, n+1)
    trees[0], trees[1] = 1, 1
    for i := 2; i < n+1; i++ {
        for j := 1; j <= i; j++ {
            trees[i] = trees[i] + trees[j-1]*trees[i-j]
        }
    }
    return trees[n]
}
```

## 实现
找到规律，使用dp算法
```golang
func numTrees(n int) int {
    dp := make([]int, n)
    cal(&dp, n)
    return dp[n-1]
}

func cal(d *[]int, n int) {
    if n == 1 {
        (*d)[n-1] = 1
    } else {
        cal(d, n-1)
        (*d)[n-1] = (4*n - 2) * (*d)[n-2] / (n + 1)
    }
}
```

## 改进
```golang

```

## 反思

## 参考