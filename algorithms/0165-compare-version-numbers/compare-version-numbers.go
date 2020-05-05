package compare_version_numbers

func compareVersion(version1 string, version2 string) int {
	var num1, num2 int
	var i, j int
	len1, len2 := len(version1), len(version2)
	for {
		// get each part of version1
		for ; i < len1; i++ {
			if version1[i] == '.' {
				break
			}
			num1 += 10*num1 + int(version1[i]-'0')
		}
		i++

		// get each part of version2
		for ; j < len2; j++ {
			if version2[j] == '.' {
				break
			}
			// cal sum
			num2 += 10*num2 + int(version2[j]-'0')
		}
		j++

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		} else {
			if j >= len2 && i >= len1 {
				break
			}
			num1, num2 = 0, 0
		}
	}
	return 0
}
