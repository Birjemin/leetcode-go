package gas_station

func canCompleteCircuit(gas []int, cost []int) int {

    length := len(gas)

    // get subtraction
    for i := range gas {
        gas[i] -= cost[i]
    }

    for i, v := range gas {
        // next value
        if v < 0 {
            continue
        }

        for j := 1; j <= length; j++ {
            // got it
            if j == length {
                return i
            }

            p := j + i
            if p >= length {
                p -= length
            }

            v += gas[p]
            if v < 0 {
                break
            }
        }
    }

    return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {
    // now gas
    // left gas
    var now, left, p int
    for i, v := range gas {
        now += v - cost[i]
        left += v - cost[i]

        if now < 0 {
            now = 0
            p = i
        }
    }
    if left >= 0 {
        return p + 1
    }
    return -1

}
