package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Message struct to be sent as JSON to SNS
type Message struct {
	Text string `json:"text"`
}

// Lambda handler function
func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Specify the SNS topic ARN
	snsTopicARN := "arn:aws:sns:eu-central-1:996985152674:mytopic"

	// Create an SNS session
	sess := session.Must(session.NewSession())
	snsClient := sns.New(sess)

	// Create a message
	message := Message{
		Text: "Hello from Lambda!",
	}

	// Convert message to JSON
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling JSON"}, err
	}

	// Publish message to SNS topic
	publishInput := &sns.PublishInput{
		TopicArn: aws.String(snsTopicARN),
		Message:  aws.String(string(messageJSON)),
	}

	_, err = snsClient.Publish(publishInput)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error publishing message to SNS"}, err
	}

	// Return success response
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Message sent to SNS topic",
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
