package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
}

func SendMessageToSQS(ctx context.Context) error {
	// Create a session and SQS client
	sess := session.Must(session.NewSession())
	sqsSvc := sqs.New(sess)

	// Define the message content you want to send
	messageBody := "This is the message content."

	// Specify the SQS queue URL where you want to send the message
	queueURL := "https://sqs.eu-central-1.amazonaws.com/996985152674/article"

	// Send the message to the SQS queue
	_, err := sqsSvc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	})

	if err != nil {
		logger.Error("failed to send message to SQS", zap.Error(err))
		return err
	}

	logger.Info("message sent to SQS successfully")
	return nil
}

func main() {
	lambda.Start(SendMessageToSQS)
}
