package set_matrix_zeroes

func setZeroes(matrix [][]int) {
    height := len(matrix)
    if height == 0 {
        return
    }
    width := len(matrix[0])
    var h, w bool
    // get status of first row and first col
    for i := 0; i < width; i++ {
        if matrix[0][i] == 0 {
            w = true
        }
    }
    for i := 0; i < height; i++ {
        if matrix[i][0] == 0 {
            h = true
        }
    }

    // move
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {

            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }

    // set col
    for i := 1; i < width; i++ {
        if matrix[0][i] == 0 {
            for j := 1; j < height; j++ {
                matrix[j][i] = 0
            }
        }
    }

    // set row
    for i := 1; i < height; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < width; j++ {
                matrix[i][j] = 0
            }
        }
    }

    // set first col
    if h {
        for i := 0; i < height; i++ {
            matrix[i][0] = 0
        }

    }

    // set first row
    if w {
        for i := 0; i < width; i++ {
            matrix[0][i] = 0
        }
    }

    return
}