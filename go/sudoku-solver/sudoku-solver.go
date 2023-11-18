func solveSudoku(board [][]byte)  {
    solve(board)
}

func solve(board [][]byte) bool {
    for i := 0; i < len(board); i ++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == byte('.') {
                for c := byte('1'); c <= byte('9'); c++ {
                    if isValid(board, i, j, c) {
                        board[i][j] = c
                        
                        if solved := solve(board); solved {
                            return true
                        } else {
                            board[i][j] = byte('.')
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
    for i := 0; i < 9; i++ {
        if board[i][col] == c ||board[row][i] == c || board[3 * (row/3) + i / 3][3 * (col / 3) + i % 3] == c{
            return false
        }
    }
    return true
}
