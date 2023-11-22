class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        valid = True
        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    continue
                if not isValid(board, i, j):
                    return False
        return valid

def isValid(board, row, col):
    b = board[row][col]
    for i in range(9):
        r = 3 * (row // 3) + i // 3
        c = 3 * (col // 3) + i % 3
        if (board[row][i] == b and col != i) or (board[i][col] == b and row != i) or (board[r][c] == b and row != r and col != c):
            return False
    return True
