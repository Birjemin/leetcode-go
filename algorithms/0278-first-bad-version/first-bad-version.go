package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var badBlock int

func setBadVersion(version int) {
	badBlock = version
}

func isBadVersion(version int) bool {
	return badBlock == version
}

func firstBadVersion(n int) int {
	return cal(1, n)
}

func cal(min, max int) int {
	if min == max {
		return min
	}
	i := (min + max) / 2
	if isBadVersion(i) {
		return cal(min, i)
	} else {
		return cal(i+1, max)
	}
}

func firstBadVersion1(n int) int {

	for l := 1; l != n; {
		mid := (l + n) / 2
		if isBadVersion(mid) {
			n = mid
		} else {
			l = mid + 1
		}
	}
	return n
}
