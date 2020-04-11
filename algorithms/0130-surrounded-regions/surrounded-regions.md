## 问题
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:
```
X X X X
X O O X
X X O X
X O X X
```

After running your function, the board should be:
```
X X X X
X X X X
X X X X
X O X X
```

Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

## 分析
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
反向思维，找出入口，然后dfs查找到所有没有闭合的，然后遍历闭合即可

## 最高分
```golang

```


## 实现
dfs方式，从四周找入口，然后置为A，遍历将O置为X，A置回O即可
```golang
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

```

## 改进
```golang

```

## 反思

## 参考