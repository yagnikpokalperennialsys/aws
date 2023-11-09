package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	for _, message := range sqsEvent.Records {
		logger.Info("received sqs event", zap.Any("message", message))

		// Decode JSON
		// This will give intentional error to check the dead letter queue
		event := &Event{}
		err := json.Unmarshal([]byte(message.Body), event)
		if err != nil {
			return err
		}

		logger.Info("decoded event", zap.Any("event", event))
		// Not deleting the message to retry 3 times and add the message to dead letter queue
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
