package ai

import (
	"context"
	"log"

	gemini "google.golang.org/genai"
)

func GetAIClient(ctx context.Context) *gemini.Client {
	clientConfig := gemini.ClientConfig{
		APIKey: "AIzaSyBk_YBJx4bChZHPZ_EsDXb4DxOzVbXe1rs",
	}
	client, err := gemini.NewClient(ctx, &clientConfig)

	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	return client
}
