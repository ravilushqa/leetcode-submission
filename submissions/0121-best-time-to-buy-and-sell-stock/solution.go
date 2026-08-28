func maxProfit(prices []int) int {
    minPrice := -1
    maxProfit := 0
    
    for _, v := range prices {
        if minPrice == -1 || v < minPrice {
            minPrice = v
        } else {
            if v - minPrice > maxProfit {
                maxProfit = v - minPrice
            }
        }
    }

    return maxProfit
}
