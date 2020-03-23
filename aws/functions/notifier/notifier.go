package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		log.Printf("The message %s for the source %s = %s\n", message.MessageId, message.EventSource, message.Body)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
