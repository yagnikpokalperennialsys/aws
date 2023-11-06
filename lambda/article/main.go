package main

import (
	"pet/internal/routes"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(routes.Routes)
}
