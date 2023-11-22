var isValidSudoku = function(board) {
    let valid = true
    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            if (board[i][j] == ".") {
                continue
            }
            if (!isValid(board, i, j)){
                return false
            }
        } 
    }
    return valid
};

function isValid(board, row, col){
    b = board[row][col]
    for (i = 0; i < 9; i++) {
        r = 3 * Math.floor(row / 3) + Math.floor(i / 3);
        c = 3 * Math.floor(col / 3) + Math.floor(i % 3)
        if ((board[row][i] == b && col != i) || (board[i][col] == b && row != i) || (board[r][c] == b && row != r && col != c)) {
            
            return false
        }
    }
    return true
}
