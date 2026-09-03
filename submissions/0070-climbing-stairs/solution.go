func climbStairs(n int) int {
    if n <= 3 {
        return n
    }


    prev := 2
    curr := 3
    

    for i := 4; i <= n; i++ {
        prev, curr = curr, prev + curr
    }

    return curr
}
