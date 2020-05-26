package bulls_and_cows

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	var count, ret int
	length := len(secret)
	tmp := make(map[int32]int, length)

	for i, v := range secret {
		if v == int32(guess[i]) {
			count++
		}
		if _, ok := tmp[v]; ok {
			tmp[v]++
		} else {
			tmp[v] = 1
		}
	}
	for _, v := range guess {
		if _, ok := tmp[v]; ok {
			if tmp[v] > 0 {
				ret++
			}
			tmp[v]--
		}
	}
	return fmt.Sprintf("%dA%dB", count, ret-count)
}

func getHint1(secret string, guess string) string {
	var arr1 [10]int
	var count, ret int
	length := len(secret)

	for i := 0; i < length; i++ {
		s1, s2 := secret[i]-'0', guess[i]-'0'
		if s1 == s2 {
			count++
		} else {
			arr1[s1]++
			arr1[s2]--
		}
	}

	for _, v := range arr1 {
		if v > 0 {
			ret += v
		}
	}
	return fmt.Sprintf("%dA%dB", count, length-count-ret)
}
