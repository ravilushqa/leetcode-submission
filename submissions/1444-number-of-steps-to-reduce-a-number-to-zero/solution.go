func numberOfSteps(num int) int {
    var res int
    
    for num > 0 {
        if num % 2 == 1 {
            num--
        } else {
            num/=2
        }
        res++
    }
    
    return res
}
