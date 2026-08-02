import (
    "fmt"
)

func summaryRanges(nums []int) []string {
    var res []string
    var i int

    for i < len(nums) {
        start := nums[i]
        for i < len(nums) - 1 && nums[i]+1 == nums[i+1] {
            i++
            continue
        }

        if start == nums[i] {
            res = append(res, fmt.Sprintf("%d", start))
            i++
            continue
        }

        res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))        
        i++
    }

    return res
}
