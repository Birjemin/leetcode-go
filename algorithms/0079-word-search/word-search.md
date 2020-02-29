## 问题
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:
```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
```

## 分析
在矩阵中查找一个字符串是否存在 dfs方式

## 最高分
```golang
func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if DFS(board, i, j, word, 0) {
                return true
            }
        }
    }
    return false
}

func DFS(board [][]byte, i int, j int, word string, index int) bool {
    if index == len(word) {
        return true
    }
    
    if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[index] != board[i][j] {
        return false
    }
    
    temp := board[i][j]
    board[i][j] = '%'
    
    a := DFS(board, i + 1, j, word, index + 1) || DFS(board, i, j+1, word, index + 1) || DFS(board, i-1, j, word, index+1) || DFS(board, i, j-1, word, index+1)
    
    board[i][j] = temp
    return a
}

```

## 实现
使用dfs搜索，先找到target，然后搜索它的上下左右值
- 注意判断矩阵边界
- 是否重复，比如从(1,1) -> (1,2)，那么(1,2)只有三个方向可以搜索(1,3)(2,2)(0,2)，不可以搜索(1,1)
```golang
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

```

## 改进
```golang

```

## 反思

## 参考