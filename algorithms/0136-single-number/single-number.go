package single_number

func singleNumber(nums []int) int {
    ret := make(map[int]bool, (len(nums)+1)/2)
    for _, v := range nums {
        if ret[v] {
            ret[v] = false
        } else {
            ret[v] = true
        }
    }
    for k, v := range ret {
        if v {
            return k
        }
    }
    return 0
}

func singleNumber1(nums []int) int {
    length := len(nums)
    for i := 1; i < length; i++ {
        nums[0] ^= nums[i]
    }
    return nums[0]
}
