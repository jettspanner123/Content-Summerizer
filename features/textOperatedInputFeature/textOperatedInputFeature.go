package textOperatedInputFeature

import (
	"context"
	ai "csum/ai"
	helpers "csum/helpers/functions"
	"fmt"
	"log"

	gemini "google.golang.org/genai"
)

func TextOperatedInputFeature(text string, outputLength string) {

	// MARK: Creating Gemini Client
	ctx := context.Background()
	geminiClient := ai.GetAIClient(ctx)

	// MARK: Get Output Length
	outputLengthT, err := helpers.GetOutputLength(outputLength)
	if err != nil {
		log.Fatal("Error converting output length option to integer")
	}

	// MARK: System Prompt

	systemPrompt := fmt.Sprintf(`
		YOUR ROLE:
			- You are a ai, that generate text based on the topic, with a certain length.
		
		YOUR TOPIC: 
			- %s is your topic to generate.
			- %d words, this is the length in which you want to generate.
		
		RULES TO FOLLOW:
			- No matter what the input text is, you have to keep it in the range of %d words.
	`, text, outputLengthT, outputLengthT,
	)
	result, err := geminiClient.Models.GenerateContent(
		ctx, "gemini-2.5-flash", gemini.Text(systemPrompt), nil)

	if err != nil {
		log.Fatal("Error generating content: ", err)
	}

	fmt.Println(result.Text())
}
