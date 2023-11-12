package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Create a new SNS client.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	snsClient := sns.New(sess)

	// Create a message attributes map.
	messageAttributes := map[string]*sns.MessageAttributeValue{
		"eventType": {
			DataType:    aws.String("String"),
			StringValue: aws.String("sqs_specific_event"),
		},
	}

	// Create a message input.
	messageInput := &sns.PublishInput{
		MessageAttributes: messageAttributes,
		Message:           aws.String("hello beautifull world"), // Corrected field name
		TopicArn:          aws.String("arn:aws:sns:eu-central-1:996985152674:mytopic"),
	}

	// Publish the message.
	_, err := snsClient.Publish(messageInput)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Return a success response.
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Message published successfully.",
	}, nil
}
