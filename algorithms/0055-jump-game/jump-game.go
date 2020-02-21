package jump_game

func canJump(nums []int) bool {
    var k int
    for i, v := range nums {
        // if i > largest jumping distance, that mean never touch the end of array
        if i > k {
            return false
        }
        // get the largest jumping distance
        k = max(k, i+v)
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
