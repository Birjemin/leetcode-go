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
    total, curr, start := 0, 0, 0
    for i, v := range gas {
        // v - cost
        total += v - cost[i]
        curr += v - cost[i]

        if curr < 0 {
            start, curr = i+1, 0
        }
    }

    // got if
    if total >= 0 {
        return start
    }

    return -1
}
