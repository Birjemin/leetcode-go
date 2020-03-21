package majority_element

import (
    "sort"
)

func majorityElement(nums []int) int {
    m := make(map[int]int, len(nums)/2)
    for _, v := range nums {
        if _, ok := m[v]; ok {
            m[v]++
        } else {
            m[v] = 1
        }
    }
    var ret, count int
    for i, v := range m {
        if count < v {
            count, ret = v, i
        }
    }
    return ret
}

func majorityElement1(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
    var ele, count int
    for _, v := range nums {
        if count == 0 {
            ele, count = v, 1
            continue
        }
        if ele != v {
            count--
        } else {
            count++
        }
    }
    return ele
}
