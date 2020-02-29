package word_search

func exist(board [][]byte, word string) bool {
    height := len(board)
    if height == 0 {
        return false
    }
    width := len(board[0])

    // init malloc memory
    check := make([]bool, height*width)

    // check each of board
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if dfs(board, check, width, height, i, j, word, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(tmp [][]byte, check []bool, width, height, i, j int, str string, from int) bool {
    // if the element had checked
    // if the element is not equal to str[0]
    p := i*width + j
    if check[p] || tmp[i][j] != str[0] {
        return false
    }

    // end
    if str[1:] == "" {
        return true
    }

    // searching around of the element
    check[p] = true
    // top
    if i != 0 && from != 1 && dfs(tmp, check, width, height, i-1, j, str[1:], 2) {
        return true
    }

    // bottom
    if i != height-1 && from != 2 && dfs(tmp, check, width, height, i+1, j, str[1:], 1) {
        return true
    }

    // left
    if j != 0 && from != 3 && dfs(tmp, check, width, height, i, j-1, str[1:], 4) {
        return true
    }

    // right
    if j != width-1 && from != 4 && dfs(tmp, check, width, height, i, j+1, str[1:], 3) {
        return true
    }

    // if it is not be found，reverse the value of check
    check[p] = false
    return false
}
