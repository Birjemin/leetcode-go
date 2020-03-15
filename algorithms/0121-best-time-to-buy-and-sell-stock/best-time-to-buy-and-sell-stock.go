package best_time_to_buy_and_sell_stock

import (
    "math"
)

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

func maxProfit1(prices []int) int {
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