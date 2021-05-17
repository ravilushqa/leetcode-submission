func maximumWealth(accounts [][]int) int {
    var res int
    
    for _, account := range accounts {
        var accSum int
        for _, v := range account {
            accSum+=v
        }
        if accSum > res {
            res = accSum
        }
    }
    
    return res
}
