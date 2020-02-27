package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    width := len(matrix[0])
    if width == 0 {
        return false
    }
    var row int
    // find row
    for i, v := range matrix {
        if v[width-1] == target {
            return true
        } else if v[width-1] > target {
            row = i
            break
        }
    }

    // split
    left, mid := 0, width/2
    for ; mid != left; mid = (width-left)/2 + left {
        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] > target {
            width = mid
        } else {
            left = mid
        }
    }
    if matrix[row][mid] == target {
        return true
    }
    return false
}
