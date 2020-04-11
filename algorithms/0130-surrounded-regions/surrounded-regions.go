package surrounded_regions

func solve(board [][]byte) {
    height := len(board)
    if height <= 1 {
        return
    }
    width := len(board[0])

    // top and bottom
    for i := 0; i < width; i++ {
        if board[0][i] == 'O' {
            dfs(board, width, height, i, 0)
        }
        if board[height-1][i] == 'O' {
            dfs(board, width, height, i, height-1)
        }
    }

    // left and right
    for i := 1; i < height-1; i++ {
        if board[i][0] == 'O' {
            dfs(board, width, height, 0, i)
        }
        if board[i][width-1] == 'O' {
            dfs(board, width, height, width-1, i)
        }
    }

    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == 'A' {
                board[i][j] = 'O'
            }
        }
    }
}

func dfs(board [][]byte, width, height, i, j int) {
    if i < 0 || i >= width || j < 0 || j >= height || board[j][i] != 'O' {
        return
    }

    board[j][i] = 'A'
    dfs(board, width, height, i, j+1)
    dfs(board, width, height, i, j-1)
    dfs(board, width, height, i+1, j)
    dfs(board, width, height, i-1, j)
}
