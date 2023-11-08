package main

import (
	"context"
	"encoding/json"
	"time" // Import the time package

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

		// Sleep for 13 seconds
		time.Sleep(13 * time.Second)

		// Decode JSON
		event := &Event{}
		err := json.Unmarshal([]byte(message.Body), event)
		if err != nil {
			return err
		}

		logger.Info("decoded event", zap.Any("event", event))
		// Not deleting the message from the queue because we set a visibility timeout
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
