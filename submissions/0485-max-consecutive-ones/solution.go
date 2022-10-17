func findMaxConsecutiveOnes(nums []int) int {
    var cur, max int
    
    for _, v := range nums {
        if v == 1 {
            cur++
            continue
        }
        
        if cur > max {
            max = cur
        }
        cur = 0
    }
    if cur > max {
        max = cur
    }
    
    return max
}
