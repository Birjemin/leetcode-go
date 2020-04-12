package single_number_ii

func singleNumber(nums []int) int {
    ret := make(map[int]int, (len(nums)+2)/3)
    for _, v := range nums {
        if ret[v] >= 1 {
            ret[v]++
        } else {
            ret[v] = 1
        }
    }
    for k, v := range ret {
        if v == 1 {
            return k
        }
    }
    return 0
}

func singleNumber1(nums []int) int {
    ones, twos := 0, 0
    for i := 0; i < len(nums); i++ {
        ones = (ones ^ nums[i]) & ^twos
        twos = (twos ^ nums[i]) & ^ones
    }
    return ones
}
