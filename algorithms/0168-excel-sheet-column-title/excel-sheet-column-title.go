package excel_sheet_column_title

import (
    "bytes"
)

func convertToTitle(n int) string {
    ret := new(bytes.Buffer)
    for n > 26 {
        remainder := n % 26
        n /= 26
        if remainder == 0 {
            ret.WriteRune(rune(90))
            n -= 1
        } else {
            ret.WriteRune(rune(remainder + 64))
        }
    }
    ret.WriteRune(rune(n + 64))
    bytesArr := ret.Bytes()
    for from, to := 0, len(bytesArr)-1; from < to; from, to = from+1, to-1 {
        bytesArr[from], bytesArr[to] = bytesArr[to], bytesArr[from]
    }
    return string(bytesArr)
}

func convertToTitle1(n int) string {
    var ret string
    for n > 0 {
        n--
        ret = string(rune(n % 26 + 65))+ret
        n /= 26
    }
    return string(ret)
}