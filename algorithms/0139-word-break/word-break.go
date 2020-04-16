package word_break

func wordBreak(s string, wordDict []string) bool {
    // got length of wordDict and record memo
    wordDictMap, memo := make(map[string]int), make(map[string]int)
    for _, s := range wordDict {
        wordDictMap[s] = len(s)
    }

    return bfs(map[string]int{s: len(s)}, wordDictMap, &memo)
}

func bfs(str map[string]int, wordDictMap map[string]int, memo *map[string]int) bool {
    if len(str) == 0 {
        return false
    }
    queue := make(map[string]int)
    for s, sLen := range str {
        for w, wLen := range wordDictMap {

            // not fit or got it
            if sLen < wLen || s[:wLen] != w {
                continue
            }

            left := s[wLen:]

            // found!  end
            if left == "" {
                return true
            }

            // if it has in memo, just continue and it is not necessary to be stored in queue
            if _, ok := (*memo)[left]; ok {
                continue
            }

            (*memo)[left] = sLen - wLen

            // if not exist in queue, stored the value
            if _, ok := queue[left]; !ok {
                queue[left] = sLen - wLen
            }
        }
    }
    return bfs(queue, wordDictMap, memo)
}

func wordBreak1(s string, wordDict []string) bool {

    // got length of wordDict and record memo
    length := len(s)
    wordDictMap, dp := make(map[string]int, length), make(map[int]bool, length+1)
    for _, s := range wordDict {
        wordDictMap[s] = len(s)
    }
    dp[0] = true
    for i := 1; i <= length; i++ {
        for j := 0; j < i; j++ {
            if !dp[j] {
                continue
            }
            if _, ok := wordDictMap[s[j:i]]; ok {
                dp[i] = true
                break
            }
        }
    }
    return dp[length]
}
