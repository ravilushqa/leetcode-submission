func summaryRanges(nums []int) []string {
    res := make([]string, 0)
    i := 0

    for i < len(nums) {
        start := nums[i]

        for i + 1 < len(nums) && nums[i] + 1 == nums[i + 1] {
            i++
            continue
        }

        if start == nums[i] {
            res = append(res, fmt.Sprintf("%d", start))
        } else {
            res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
        }
        
        i++
    }

    return res
}
