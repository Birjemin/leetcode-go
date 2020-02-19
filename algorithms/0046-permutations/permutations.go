package permutations

func permute(nums []int) [][]int {
    length := len(nums)
    if length == 1 {
        return [][]int{
            nums,
        }
    }
    var ret [][]int
    cal(&ret, nums, []int{})
    return ret
}

func cal(ret *[][]int, nums, temp []int) {
    if len(nums) == 0 {
        b := make([]int, len(temp))
        copy(b, temp)
        *ret = append(*ret, b)
    }
    for i, v := range nums {
        if i == 0 {
            cal(ret, nums[i+1:], append(temp, v))
        } else {
            b := make([]int, len(nums))
            copy(b, nums)
            cal(ret, append(b[:i], b[i+1:]...), append(temp, v))
        }
    }
}
