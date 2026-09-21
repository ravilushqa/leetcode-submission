func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    

    for i, v := range nums {
        if idx, found := m[target-v]; found {
            return []int{idx, i}
        }

        m[v] = i
    }

    return nil
}
