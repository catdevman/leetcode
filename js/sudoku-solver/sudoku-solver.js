var solveSudoku = function(board) {
    solve(board);
};

function isValid(board, row, col, num) {
    for (let i = 0; i < 9; i++) {
        if (board[row][i] === num || board[i][col] === num ||
            board[3 * Math.floor(row / 3) + Math.floor(i / 3)][3 * Math.floor(col / 3) + i % 3] === num) {
            return false;
        }
    }
    return true;
}

function solve(board) {
    for (let row = 0; row < 9; row++) {
        for (let col = 0; col < 9; col++) {
            if (board[row][col] === ".") {
                for (let num = 1; num <= 9; num++) {
                    if (isValid(board, row, col, num.toString())) {
                        board[row][col] = num.toString();
                        if (solve(board)) {
                            return true;
                        }
                        board[row][col] = ".";
                    }
                }
                return false;
            }
        }
    }
    return true;
}
