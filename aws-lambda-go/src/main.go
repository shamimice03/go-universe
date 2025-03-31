package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// handleRequest processes the Lambda event
func handleRequest(ctx context.Context, event any) (string, error) {
	// First, capture the event as a single string variable
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return "", err
	}

	eventStr := string(eventBytes)

	// Log the entire event as a single message - CloudWatch will preserve this formatting
	// The key is that we're making a single call to log.Print to keep it as one entry
	message := fmt.Sprintf("EVENT_DATA_BEGIN\n%s\nEVENT_DATA_END", eventStr)
	log.Print(message)

	return "Event processed successfully", nil
}

func main() {
	lambda.Start(handleRequest)
}
