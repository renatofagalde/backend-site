package main

import (
	secrets "backend-site/src/config/aws.secrets"
	"backend-site/src/config/database/mysqldb"
	"backend-site/src/controller/routes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
)

var ginLambda *ginadapter.GinLambda

func init() {

	log.Println("Initialize lambda")

	/*	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/
	mysqlProperties, err := secrets.GetSecrets()
	if err != nil {
		log.Fatalf("Error to get secrets")
		return
	}

	database, err := mysqldb.NewMySQLGORMConnection(context.Background(), *mysqlProperties)
	if err != nil {
		log.Fatalf("Error ao conectar no no banco, error=%s", err.Error())
		return
	} else {
		fmt.Println("conexao com sucesso")
	}

	siteController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, siteController)
	ginLambda = ginadapter.New(router)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, req)
}
