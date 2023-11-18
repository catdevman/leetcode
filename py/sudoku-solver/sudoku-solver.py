import math
class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        solve(board)

def solve(board):
    for i in range(9):
        for j in range(9):
            if board[i][j] == ".":
                for num in map(str, range(1, 10)):
                    if isValid(board, i, j, num):
                        board[i][j] = num
                        if solve(board):
                            return True
                        board[i][j] = "."
                return False
    return True

def isValid(board, row, col, num):
    # Check if the number is not in the current row, column and 3x3 sub-box
    for i in range(9):
        if board[row][i] == num or board[i][col] == num or \
           board[3 * (row // 3) + i // 3][3 * (col // 3) + i % 3] == num:
            return False
    return True
