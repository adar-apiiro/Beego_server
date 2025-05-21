package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/simongottschlag/azure-openai-gpt-slack-bot/internal/gpt3"
)

func main() {
	// Retrieve necessary environment variables
	apiKey := os.Getenv("AZURE_OPENAI_API_KEY")
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	deployment := os.Getenv("AZURE_OPENAI_DEPLOYMENT")

	// Validate environment variables
	if apiKey == "" || endpoint == "" || deployment == "" {
		log.Fatal("Please set AZURE_OPENAI_API_KEY, AZURE_OPENAI_ENDPOINT, and AZURE_OPENAI_DEPLOYMENT environment variables.")
	}

	// Initialize the GPT-3 client
	client := gpt3.NewClient(apiKey, endpoint, deployment)

	// Define the prompt
	prompt := "Hello, how can I assist you today?"

	// Send the prompt to the GPT-3 service
	response, err := client.Complete(context.Background(), prompt)
	if err != nil {
		log.Fatalf("Error completing prompt: %v", err)
	}

	// Output the response
	fmt.Println("GPT-3 Response:", response)
}
