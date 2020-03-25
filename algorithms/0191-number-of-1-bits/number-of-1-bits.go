package number_of_1_bits

func hammingWeight(num uint32) int {
    var count int
    for i := 0; i < 32; i++ {
        if num&1 == 1 {
            count++
        }
        num >>= 1
    }
    return count
}

func hammingWeight1(num uint32) int {
    var count int
    for ; num >= 1; num >>= 1 {
        if num&1 == 1 {
            count++
        }
    }
    return count
}
