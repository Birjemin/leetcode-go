package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    min, max := prices[0], 0
    for _, v := range prices {
        if v < min {
            min = v
        } else if v-min > 0 {
            max += v - min
            min = v
        }
    }
    return max
}

func maxProfit1(prices []int) int {
    length := len(prices)
    var max int
    for i := 1; i < length; i++ {
        if prices[i] > prices[i-1] {
            max += prices[i] - prices[i-1]
        }
    }
    return max
}
