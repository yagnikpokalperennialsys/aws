package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync() // flushes buffer, if any
}

type Event struct {
	Name string `json:"name"`
}

func MyHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// Create a session and SQS client
	sess := session.Must(session.NewSession())
	sqsSvc := sqs.New(sess)

	for _, message := range sqsEvent.Records {
		//logger.Info("received sqs event", zap.Any("message", message.))
		fmt.Println(message.Body)
		// Delete the message from the SQS queue refer https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/step-receive-delete-message.html
		sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"), // Replace with your SQS queue URL
			ReceiptHandle: aws.String(message.ReceiptHandle),
		})

	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
