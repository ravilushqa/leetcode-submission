func flipAndInvertImage(image [][]int) [][]int {
    for i, v := range image {
        image[i] = invert(flip(v))
    }
    return image
}

func flip(image []int) []int {
    for i := range image {
        if i >= len(image)/2  {
            break
        }
        
        image[i], image[len(image) - 1 - i] = image[len(image) - 1 -i], image[i]
    }
     return image
}

func invert(image []int) []int {
    for i := range image {
        if image[i] == 0 {
            image[i] = 1
            continue
        }
        image[i] = 0
    }
                                                
    return image
}
