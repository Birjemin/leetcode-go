## 问题
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:
```
Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
```

Example 2:
```
Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
```
Explanation: Same as Example 1, except with the 5 in the top left corner being 
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.

## 分析
判断一个数独是否有效
- 横向没有重复
- 纵向没有重复
- 方格没有重复

## 最高分
```golang
func isValidSudoku(board [][]byte) bool {
    var row, col, block [10]int
    for i := range board {
        for j, c := range board[i] {
            if c == '.' {
                continue
            }
            v := 1 << int(c-'0')
            if (col[j]&v)|(row[i]&v)|(block[i/3+j/3*3]&v) > 0 {
                return false
            }
            col[j] |= v
            row[i] |= v
            block[i/3+j/3*3] |= v
        }
    }
    return true
}
```

## 实现
其实就是判断是否存在的问题
```golang
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
```

## 改进
改成位运算，提交的记录来看好像也就那样
```golang
func isValidSudoku(board [][]byte) bool {
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
```

## 反思
可以看一下go关于锁的实现，大量的位运算~~~刘辟！

## 参考
