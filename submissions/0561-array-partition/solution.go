func arrayPairSum(nums []int) int {
    var res int
    sort.Ints(nums)
    for i:=0;i<len(nums)-1;i+=2 {
        res += nums[i]
    }
    
    return res
}
