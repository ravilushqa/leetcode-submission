func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := findMax(candies)
    output := make([]bool, len(candies))
    for i, v := range candies {
        output[i] = v + extraCandies >= max
    }
    return output
}

func findMax(candies []int) int {
    max := 1 // Because 1 <= candies[i] <= 100
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    return max
}
