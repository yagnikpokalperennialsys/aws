package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

type SNSMessage struct {
	Message string `json:"Message"`
}

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	config.DisableStacktrace = true

	var err error
	logger, err = config.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()
}

func handleSNSMessage(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		rawMessage := record.SNS.Message
		logger.Info("Received SNS message", zap.String("message", rawMessage))
	}

	return nil
}

func main() {
	lambda.Start(handleSNSMessage)
}
