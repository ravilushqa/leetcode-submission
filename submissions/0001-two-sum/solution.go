func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i,v := range nums {
        if j, exist := m[target-v]; exist {
            return []int{j,i}
        }

        m[v] = i
    }

    return nil
}
