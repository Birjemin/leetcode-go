package spiral_matrix_ii

func generateMatrix(n int) [][]int {
    // init the array
    ret, v := make([][]int, n), 1
    for i := range ret {
        ret[i] = make([]int, n)
    }

    // the number of circle
    for i := 0; i < n; i++ {
        // moving times of each circle
        row := n - i
        // first
        for a := i; a < row; a++ {
            ret[i][a] = v
            v++
        }
        // second
        for b := i + 1; b < row; b++ {
            ret[b][row-1] = v
            v++
        }

        // three
        for c := row - 2; c > i; c-- {
            ret[row-1][c] = v
            v++
        }

        // four
        for d := row - 1; d > i; d-- {
            ret[d][i] = v
            v++
        }
    }
    return ret
}
