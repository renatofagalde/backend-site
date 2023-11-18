package main

import (
	"backend-site/src/controller/routes"
	controller "backend-site/src/controller/site"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, controller.NewSiteControllerInterface())
	ginLambda = ginadapter.New(router)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, req)
}
