package group_anagrams

import (
    "sort"
)

type bytes []byte

func (b bytes) Len() int {
    return len(b)
}

func (b bytes) Less(i, j int) bool {
    return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
    var ret [][]string
    var j int
    temp := make(map[string]int)
    length := len(strs)
    for i := 0; i < length; i++ {
        key := bytes(strs[i])
        sort.Sort(key)
        str := string(key)
        // exist
        if _, ok := temp[str]; ok {
            ret[temp[str]] = append(ret[temp[str]], strs[i])
        } else {
            temp[str] = j
            ret = append(ret, []string{strs[i]})
            j++
        }
    }
    return ret
}

func groupAnagrams1(strs []string) [][]string {
    var ret [][]string
    hm := make(map[string][]string)
    for _, str := range strs {
        key := bytes(str)
        sort.Sort(key)
        str1 := string(key)
        hm[str1] = append(hm[str1], str)
    }
    for _, words := range hm {
        ret = append(ret, words)
    }
    return ret
}
