package remove_element

func removeElement1(nums []int, val int) int {
    left, flag, num, length := 0, -1, 0, len(nums)
    for ; left < length; left++ {
        if nums[left] == val {
            if flag == -1 {
                flag = left
            }
        } else {
            if flag != -1 {
                nums[flag], nums[left] = nums[left], nums[flag]
                left = flag
                flag = -1
            }
            num++
        }
    }
    return num
}

func removeElement(nums []int, val int) int {
    var l int
    for _, v := range nums {
        if v != val {
            nums[l] = v
            l++
        }
    }
    return l
}
