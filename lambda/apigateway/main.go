package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type TeamRequest struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Response struct {
	Message string `json:"message"`
}

func GetTeamAndMembers(req TeamRequest) (Response, error) {
	if req.FirstName == "" || req.Lastname == "" {
		return Response{
			Message: "Error in firstname or lastname",
		}, nil
	}
	return Response{
		Message: "Hello " + req.FirstName + " " + req.Lastname,
	}, nil
}

func main() {
	lambda.Start(GetTeamAndMembers)
}
