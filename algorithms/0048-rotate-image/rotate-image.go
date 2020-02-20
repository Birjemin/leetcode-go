package rotate_image

func rotate(matrix [][]int) {
    length := len(matrix)
    // the number of circle
    for i := 0; i < length/2; i++ {
        // moving times of each circle
        for j := i; j < length-i-1; j++ {
            // temp
            temp := matrix[i][j]
            // one
            matrix[i][j] = matrix[length-j-1][i]
            // two
            matrix[length-j-1][i] = matrix[length-i-1][length-j-1]
            // three
            matrix[length-i-1][length-j-1] = matrix[j][length-i-1]
            // four
            matrix[j][length-i-1] = temp
        }
    }
}

func rotate1(matrix [][]int) {
    length := len(matrix)
    // the number of circle
    for i := 0; i < length/2; i++ {
        // moving times of each circle
        for j := i; j < length-i-1; j++ {
            matrix[i][j], matrix[length-j-1][i], matrix[length-i-1][length-j-1], matrix[j][length-i-1] = matrix[length-j-1][i], matrix[length-i-1][length-j-1], matrix[j][length-i-1], matrix[i][j]
        }
    }
}
