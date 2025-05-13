package middlewares

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
)

func APIGatewayContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("X-Golambdaproxy-Apigw-Context")
		if val == "" {
			c.Next()
			return
		}

		var apiGatewayContext events.APIGatewayProxyRequestContext
		if err := json.Unmarshal([]byte(val), &apiGatewayContext); err != nil {
			log.Println("Error unmarshalling API Gateway context:", err)
			c.Next()
			return
		}

		c.Set("apiGatewayContext", apiGatewayContext)

		c.Next()
	}
}

func GetAPIGatewayEvent(c *gin.Context) (events.APIGatewayProxyRequestContext, bool) {
	event, exists := c.Get("apiGatewayContext")
	if !exists {
		return events.APIGatewayProxyRequestContext{}, false
	}

	return event.(events.APIGatewayProxyRequestContext), true
}
