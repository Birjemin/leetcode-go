package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
    slow, fast := 1, len(numbers)
    for slow < fast {
        switch sum := numbers[slow-1] + numbers[fast-1] - target; {
        case sum > 0:
            fast -= 1
        case sum < 0:
            slow += 1
        default:
            return []int{slow, fast}
        }
    }
    return []int{}
}
