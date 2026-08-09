func maxDistToClosest(seats []int) int {
    var res int
    prevSeat := -1

    for i, v  := range seats {
        if v == 1 {
            if prevSeat == -1 {
                res = i
            } else {
                k := i - prevSeat - 1
                res = max(res, (k + 1) / 2 )
            }
            
            prevSeat = i
        }
    }

    res = max(res, len(seats) - 1 - prevSeat)

    return res
}
