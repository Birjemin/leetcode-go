package valid_sudoku

func isValidSudoku(board [][]byte) bool {
    rowMask, colMask, areaMask := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
    var v byte
    var idx int
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            // break if element is valid
            if board[i][j] == '.' {
                continue
            }

            // row
            v = board[i][j] - '0' - 1
            if rowMask[i][v] {
                return false
            }
            rowMask[i][v] = true

            // col
            if colMask[j][v] {
                return false
            }
            colMask[j][v] = true

            // box
            idx = (i/3)*3 + j/3
            if areaMask[idx][v] {
                return false
            }
            areaMask[idx][v] = true
        }
    }
    return true
}

func isValidSudoku1(board [][]byte) bool {
    rowMask, colMask, areaMask := [9]uint{}, [9]uint{}, [9]uint{}
    var v, idx uint
    for i := range board {
        for j := range board[i] {
            // break if element is valid
            if board[i][j] == '.' {
                continue
            }

            // row and col
            v = 1 << (uint)(board[i][j]-'0'-1)
            if rowMask[i]&v != 0 || colMask[j]&v != 0 {
                return false
            }

            // box
            idx = (uint)((i/3)*3 + j/3)
            if areaMask[idx]&v != 0 {
                return false
            }

            // handle value
            rowMask[i] |= v
            colMask[j] |= v
            areaMask[idx] |= v
        }
    }
    return true
}
