package routes

import (
	"context"
	"pet/internal/controller"

	"github.com/aws/aws-lambda-go/events"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

type Router interface {
	Routes(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func Routes(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		r := gin.Default()
		r.GET("/articles", controller.GetArticles)
		r.POST("/articles/:id", controller.GetArticle)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}
