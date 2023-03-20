func canPlaceFlowers(flowerbed []int, n int) bool {
    var allowed int
    for i, v := range flowerbed {
        if v == 1 {
            continue
        }

        // v == 0
        if (i == 0 || flowerbed[i-1] == 0) &&  (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
            allowed++
            flowerbed[i] = 1
        }
    }

    return n <= allowed
}
