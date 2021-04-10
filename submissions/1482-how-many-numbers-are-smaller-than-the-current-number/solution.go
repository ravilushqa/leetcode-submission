func smallerNumbersThanCurrent(nums []int) []int {
    res := make([]int, len(nums))
    
    for i,v := range nums {
        for _,k := range nums {
            if v > k {
                res[i]++
            }
        }
    }
    
    return res
}
