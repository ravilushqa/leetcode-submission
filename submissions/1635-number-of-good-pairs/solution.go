func numIdenticalPairs(nums []int) int {
    var score int
    
    for i := range nums {
        for j:=i+1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                score++
            }
        }
    }
    
    return score
}
