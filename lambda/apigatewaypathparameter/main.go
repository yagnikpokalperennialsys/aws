package main

import (
	"context"
	"encoding/json"
	"strings"

	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

var heroData = map[string]string{
	"1": "Iron Man",
	"2": "Superman",
	"3": "Spider-Man",
	"4": "Black Panther",
	"5": "Wonder Woman",
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.PathParameters["proxy"]
	heroID := request.QueryStringParameters["sort"]

	if heroID != "" {
		// If 'sort' parameter is present, return a sorted list of heroes
		heroes := make([]string, 0, len(heroData))
		for _, name := range heroData {
			heroes = append(heroes, name)
		}
		sort.Strings(heroes)
		heroList := strings.Join(heroes, ", ")
		response := Response{Message: heroList}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(jsonResponse),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	hero, found := heroData[path]
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}

	response := Response{Message: hero}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
