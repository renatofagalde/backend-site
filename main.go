package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/goombaio/namegenerator"
	"net/http"
	"time"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := fmt.Sprintf("Random-> %s", namegenerator.NewNameGenerator(seed).Generate())

	return events.APIGatewayProxyResponse{
		Body:       string(nameGenerator),
		StatusCode: http.StatusOK,
	}, nil

}
