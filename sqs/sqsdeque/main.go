package main

import (
	"context"
	"encoding/json"

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
		logger.Info("received sqs event", zap.Any("message", message))

		// Decode JSON
		event := &Event{}
		err := json.Unmarshal([]byte(message.Body), event)
		if err != nil {
			return err
		}

		logger.Info("decoded event", zap.Any("event", event))

		// Delete the message from the SQS queue
		_, err = sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"), // Replace with your SQS queue URL
			ReceiptHandle: aws.String(message.ReceiptHandle),
		})

		if err != nil {
			logger.Error("failed to delete message", zap.Error(err))
		}
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
