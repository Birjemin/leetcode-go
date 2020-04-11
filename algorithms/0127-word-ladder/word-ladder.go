package word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {

    count := 1

    length := len(endWord)
    regList := getRegWordList(append(wordList, beginWord), length)

    // cal
    beginWords := []string{beginWord}
    wordLists := [][]string{wordList}
    for {
        if len(beginWords) == 0 {
            return 0
        }
        var queues []string
        var extra [][]string

        for i, b := range beginWords {
            if endWord == b {
                return count
            }
            for j, w := range wordLists[i] {

                if !dislocate(b, w, regList) {
                    continue
                }
                // find
                queues = append(queues, w)
                wordLis := delElement(j, wordLists[i])
                extra = append(extra, wordLis)
            }
        }
        beginWords = queues
        count++
        wordLists = extra
    }
}

// delete element
func delElement(idx int, wordList []string) []string {
    length := len(wordList)
    b := make([]string, length)
    copy(b, wordList)
    copy(b[idx:], b[idx+1:])
    return b[:length-1]
}

func getRegWordList(wordList []string, length int) map[string][]string {
    ret := map[string][]string{}
    for _, w := range wordList {
        ret[w] = getRegList(w, length)
    }
    return ret
}

func getRegList(s string, length int) []string {
    var ret []string
    for i := 0; i < length; i++ {
        ret = append(ret, s[0:i]+"*"+s[i+1:])
    }
    return ret
}

// find the element
func dislocate(s1, s2 string, m map[string][]string) bool {
    for _, s1 := range m[s1] {
        for _, s2 := range m[s2] {
            if s1 == s2 {
                return true
            }
        }
    }
    return false
}
