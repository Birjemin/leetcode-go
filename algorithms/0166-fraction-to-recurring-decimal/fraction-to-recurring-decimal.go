package fraction_to_recurring_decimal

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	var tmp []string
	var count int
	var flag, negative bool
	var ret string
	kv := make(map[int]int)

	// check if the number is negative
	negative = checkNegative(&numerator, &denominator)

	for {

		divider, remainder := numerator/denominator, numerator%denominator

		tmp = append(tmp, strconv.Itoa(divider))

		if remainder == 0 {
			ret = strings.Join(tmp, "")
			break
		}

		// decimal fraction
		if !flag && remainder < denominator {
			count++
			flag = true
			tmp = append(tmp, ".")
		}

		if v, ok := kv[remainder]; ok {
			ret = strings.Join(tmp[:v+1], "") + "(" + strings.Join(tmp[v+1:], "") + ")"
			break
		}

		kv[remainder] = count

		numerator = 10 * remainder
		count++
	}

	if ret != "0" && negative {
		return "-" + ret
	}
	return ret
}

// check if the number is negative
func checkNegative(numerator *int, denominator *int) (negative bool) {
	if *numerator < 0 {
		negative = true
		*numerator = -*numerator
	}

	if *denominator < 0 {
		if negative {
			negative = false
		} else {
			negative = true
		}
		*denominator = -*denominator
	}
	return
}
