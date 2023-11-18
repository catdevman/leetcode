package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type BoardEvent struct {
    Board [][]string `json:"board"`
}
type BoardResponse struct {
    Board [][]string `json:"board"`
    Solved bool `json:"solved"`
}

func Handler(ctx context.Context, event *BoardEvent) (*BoardResponse, error){
    solved := solve(event.Board)

    response := &BoardResponse{
        Board: event.Board,
        Solved: solved,
    }


    return response, nil
}

func solve(board [][]string) bool {
        for i := 0; i < len(board); i ++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == "." {
                for c := byte('1'); c <= byte('9'); c++ {
                    if isValid(board, i, j, string(c)) {
                        board[i][j] = string(c)
                        
                        if solved := solve(board); solved {
                            return true
                        } else {
                            board[i][j] = "."
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(board [][]string, row int, col int, c string) bool {
    for i := 0; i < 9; i++ {
        if board[i][col] == c ||board[row][i] == c || board[3 * (row/3) + i / 3][3 * (col / 3) + i % 3] == c{
            return false
        }
    }
    return true
}

func main(){
    lambda.Start(Handler)
}
