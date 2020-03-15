## 问题
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```

Example 2:
```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

## 分析
获取最高差，如何使用一次循环解决此类问题？

## 最高分
```golang
func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    } 
    min, maxSale := prices[0], 0
    for _,price := range prices {
        if price < min {
            min = price
        } else if (price-min) > maxSale {
            maxSale = price-min
        }
    }
    
    return  maxSale
}
```


## 实现
一次循环搞定，需要想一想
```golang
func maxProfit(prices []int) int {
    min, max := math.MaxInt64, 0
    for _, v := range prices {
        if v < min {
            min = v
        } else if v-min > max {
            max = v - min
        }
    }
    return max
}
```

## 改进
减少储存空间使用
```golang
func maxProfit(prices []int) int {
    if len(prices) ==0 {
        return 0
    }
    min, max := prices[0], 0
    for _, v := range prices {
        if v < min {
            min = v
        } else if v-min > max {
            max = v - min
        }
    }
    return max
}
```

## 反思

## 参考