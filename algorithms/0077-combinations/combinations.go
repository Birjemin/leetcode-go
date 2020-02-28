package combinations

func combine(n int, k int) [][]int {
    var ret [][]int
    dfs(&ret, []int{}, 1, n, k)
    return ret
}

func dfs(ret *[][]int, tmp []int, min, max, k int) {
    if k == 0 {
        b := make([]int, len(tmp))
        copy(b, tmp)
        *ret = append(*ret, b)
        return
    }
    // max-k+1 because the length is k
    for i := min; i <= max-k+1; i++ {
        dfs(ret, append(tmp, i), i+1, max, k-1)
    }
}
