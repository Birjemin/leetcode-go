package spiral_matrix

func spiralOrder(matrix [][]int) []int {
    col := len(matrix)
    if col == 0 {
        return []int{}
    }
    row := len(matrix[0])
    min := (min(row, col) + 1) / 2
    var ret []int
    // the number of circle
    for i := 0; i < min; i++ {
        // moving times of each circle
        row1, col1 := row-i, col-i
        // first

        for a := i; a < row1; a++ {
            ret = append(ret, matrix[i][a])
        }
        // second
        for b := i + 1; b < col1; b++ {
            ret = append(ret, matrix[b][row1-1])
        }

        // break
        if col1-i == 1 || row1-i == 1 {
            break
        }

        // three
        for c := row1 - 2; c > i; c-- {
            ret = append(ret, matrix[col1-1][c])
        }

        // four
        for d := col1 - 1; d > i; d-- {
            ret = append(ret, matrix[d][i])
        }
    }
    return ret
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}
