package main

import (
	"log"
	"net/http"
	"os"

	"github.com/InspectorGadget/ginadapter-lambda/middlewares"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

type GinResponse map[string]string

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.APIGatewayContextMiddleware())

	r.GET("/", func(c *gin.Context) {
		event, exists := middlewares.GetAPIGatewayEvent(c)
		if exists {
			c.JSON(
				http.StatusOK, GinResponse{
					"message":      "Hello from Lambda!",
					"sourceIP":     event.Identity.SourceIP,
					"awsAccountID": event.AccountID,
					"domainName":   event.DomainName,
					"stage":        event.Stage,
				},
			)
			return
		}

		c.JSON(
			http.StatusOK, GinResponse{
				"message": "Hello from Lambda! You are running the API locally.",
			},
		)
	})

	return r
}

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(event)
}

func main() {
	r := setupRouter()
	ginLambda = ginadapter.New(r)

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(Handler)
	} else {
		if err := r.Run(":3000"); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}
}
