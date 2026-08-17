func subarraysDivByK(nums []int, k int) int {
    // prefix sum
    seen := map[int]int{0:1}
    
    // is there a numbers that gives a sum 
    // v % k

    cur, res := 0, 0
    for _, v := range nums {
        cur +=v
        res += seen[Mod(cur, k)]
        seen[Mod(cur, k)]++
    }

    return res
}

func Mod(x, y int) int {
    return ((x%y) + y) % y
}
