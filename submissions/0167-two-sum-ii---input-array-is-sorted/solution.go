func twoSum(numbers []int, target int) []int {
    n := len(numbers)
    for i, v := range numbers {
        l := i + 1
        r := n - 1
        for l <= r {
            m := (l + r) / 2
            
            pair := target - v
            if pair < numbers[m] {
                r = m - 1
            } else if pair > numbers[m] {
                l = m + 1
            } else {
                return []int{i+1, m+1}
            }
        }
    }

    return []int{}
}
