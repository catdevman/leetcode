package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type BoardEvent struct {
    Board [][]string
}


func Handler(ctx context.Context, event *BoardEvent) (string, error){
    var board string

    return board, nil
}

func main(){
    lambda.Start(Handler)
}
