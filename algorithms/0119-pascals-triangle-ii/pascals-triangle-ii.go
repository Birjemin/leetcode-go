package pascals_triangle_ii

func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    ret := make([]int, rowIndex+1)
    dp(&ret, rowIndex)
    return ret
}

func dp(r *[]int, n int) {
    (*r)[n] = 1
    if n == 0 {
        return
    }
    dp(r, n-1)
    // cal the collection of n-1
    for i := n - 1; i > 0; i-- {
        (*r)[i] += (*r)[i-1]
    }
}
