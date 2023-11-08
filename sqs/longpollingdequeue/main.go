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

	// Set the visibility timeout to 10 seconds
	visibilityTimeout := int64(10)

	for _, message := range sqsEvent.Records {
		logger.Info("received sqs event", zap.Any("message", message))

		// Receive messages with long-polling and 10-second visibility timeout
		receiveParams := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"), // Replace with your SQS queue URL
			WaitTimeSeconds:     aws.Int64(10),                                                             // 10-second long-polling
			VisibilityTimeout:   aws.Int64(visibilityTimeout),
			MaxNumberOfMessages: aws.Int64(1), // Process one message at a time
		}

		result, err := sqsSvc.ReceiveMessage(receiveParams)
		if err != nil {
			logger.Error("failed to receive message with long-polling", zap.Error(err))
			return err
		}

		if len(result.Messages) > 0 {
			// Decode JSON
			event := &Event{}
			body := []byte(*result.Messages[0].Body)
			err := json.Unmarshal(body, event)
			if err != nil {
				logger.Error("failed to decode message", zap.Error(err))
				return err
			}

			logger.Info("decoded event", zap.Any("event", event))

			// Delete the message from the SQS queue
			_, err = sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"), // Replace with your SQS queue URL
				ReceiptHandle: result.Messages[0].ReceiptHandle,
			})

			if err != nil {
				logger.Error("failed to delete message", zap.Error(err))
			}
		} else {
			logger.Info("No messages received during long-polling")
		}
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
