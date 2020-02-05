package container_with_most_water

func maxArea(height []int) int {
    i, j := 0, len(height)-1
    if j == 0 {
        return 0
    }
    var max, area, h int
    // check
    for i < j {
        // left
        if height[i] < height[j] {
            h = height[i]
            i++
        } else {
            // right
            h = height[j]
            j--
        }
        //
        area = (j - i + 1) * h
        if area > max {
            max = area
        }
    }
    return max
}
