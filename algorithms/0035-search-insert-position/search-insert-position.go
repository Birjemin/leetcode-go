package search_insert_position

func searchInsert1(nums []int, target int) int {
    for k, v := range nums {
        if v >= target {
            return k
        }
    }
    return len(nums)
}

func searchInsert(nums []int, target int) int {
    var count int
    for k, v := range nums {
        if v >= target {
            return k
        }
        count++
    }
    return count
}
