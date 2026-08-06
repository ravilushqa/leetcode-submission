func longestOnes(nums []int, k int) int {
    var zeroCounter, left, res int

    for right, v := range nums {
        if v == 0 {
            zeroCounter++
        }

        for zeroCounter > k {
            if nums[left] == 0 {
                zeroCounter--
            }
            left++
        }

        if lenght := right - left + 1; lenght > res {
            res = lenght
        }
    }

    return res
}
