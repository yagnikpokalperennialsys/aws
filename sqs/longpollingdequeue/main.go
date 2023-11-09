package main

import (
	"context"
	"log"

	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func MyHandler(ctx context.Context) error {
	// Create an AWS session and SQS client
	sess := session.Must(session.NewSession())
	sqsSvc := sqs.New(sess)

	// Set the visibility timeout to 10 seconds
	visibilityTimeout := int64(10)

	// Define the time limit for polling (e.g., 15 seconds)
	startTime := time.Now()
	timeLimit := 15 * time.Second

	for {
		// Check if the time limit has been reached
		if time.Since(startTime) >= timeLimit {
			break
		}

		// Receive messages with long-polling
		receiveParams := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"),
			WaitTimeSeconds:     aws.Int64(10), // This time must be less then loop time
			VisibilityTimeout:   aws.Int64(visibilityTimeout),
			MaxNumberOfMessages: aws.Int64(10), // Adjust the number of messages to receive
		}

		result, err := sqsSvc.ReceiveMessage(receiveParams)
		if err != nil {
			return err
		}

		for _, message := range result.Messages {
			log.Printf("Processing message: %s", message.Body)

			// Delete the message from the SQS queue
			_, err := sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String("https://sqs.eu-central-1.amazonaws.com/996985152674/article"),
				ReceiptHandle: message.ReceiptHandle,
			})

			if err != nil {
				log.Printf("Failed to delete message: %v", err)
			}
		}
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
