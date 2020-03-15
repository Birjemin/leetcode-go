package pascals_triangle

func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int{}
    }
    ret := make([][]int, numRows)
    dp(&ret, numRows)

    return ret
}

func dp(r *[][]int, n int) {
    if n == 1 {
        (*r)[0] = []int{1}
        return
    }
    // cal the collection of n-1
    dp(r, n-1)

    (*r)[n-1] = []int{1}
    for i := 0; i < n-2; i++ {
        (*r)[n-1] = append((*r)[n-1], (*r)[n-2][i] + (*r)[n-2][i+1])
    }
    (*r)[n-1] = append((*r)[n-1], 1)
}
