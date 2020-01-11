package merge_sorted_array

func merge1(nums1 []int, m int, nums2 []int, n int) {
    temp := make([]int, m)
    // 将nums1复制给temp
    copy(temp, nums1)
    // j -> nums1 k -> nums2
    j, k := 0, 0
    for i := 0; i < m+n; i++ {
        // nums1用完
        if j >= m {
            nums1[i] = nums2[k]
            k++
            continue
        }
        // nums2用完
        if k >= n {
            nums1[i] = temp[j]
            j++
            continue
        }
        if temp[j] > nums2[k] {
            nums1[i] = nums2[k]
            k++
        } else {
            nums1[i] = temp[j]
            j++
        }
    }
}

func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j, k := m-1, n-1, m+n-1
    for ; i >= 0 && j >= 0; k-- {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
    // j >= 0
    for ; j >= 0; j, k = j-1, k-1 {
        nums1[k] = nums2[j]
    }
}
