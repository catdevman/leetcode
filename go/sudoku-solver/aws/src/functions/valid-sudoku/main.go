package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type BoardEvent struct {
    Board [][]string
}

type ValidationResponse struct {
    Board [][]string `json:"board"`
    Valid bool `json:"valid"`
}

func Handler(ctx context.Context, event *BoardEvent) (*ValidationResponse, error){

    response := &ValidationResponse{
        Board: event.Board,
        Valid: isValid(event.Board),
    }
    return response, nil
}

func isValid(board [][]string) bool {
    valid := true
    for i:=0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == "." {
                continue
            }
            if !currentValid(board, i, j){
                return false
            }
        } 
    }
    return valid
}

func currentValid(board [][]string, row, col int) bool{
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

func main(){
    lambda.Start(Handler)
}
