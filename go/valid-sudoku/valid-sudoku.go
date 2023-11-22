func isValidSudoku(board [][]byte) bool {
    valid := true
    for i:=0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == byte('.') {
                continue
            }
            if !isValid(board, i, j){
                return false
            }
        } 
    }
    return valid
}

func isValid(board [][]byte, row int, col int) bool{
    b := board[row][col]
    for i:=0; i < 9; i++ {
        r := 3 * (row / 3) + i / 3
        c := 3 * ( col / 3) + i % 3
        if (board[row][i] == b && col != i) || (board[i][col] == b && row != i) || (board[r][c] == b && row != r && col != c) {
            
            return false
        }
    }
    return true
}
