package textOperatedInputFeature

import (
	"context"
	ai "csum/ai"
	helpers "csum/helpers/functions"
	"fmt"
	"log"

	"github.com/fatih/color"
	gemini "google.golang.org/genai"
)

func TextOperatedInputFeature(text string, outputLength string, outputFileName *string, str string) {

	// MARK: Creating Gemini Client
	ctx := context.Background()
	geminiClient := ai.GetAIClient(ctx)

	// MARK: Get Output Length
	outputLengthT, err := helpers.GetOutputLength(outputLength)
	if err != nil {
		log.Fatal("Error converting output length option to integer")
	}

	// MARK: Validate Structure
	_, err = helpers.IsValidStructure(str)
	if err != nil {
		log.Fatal("Error validating structure")
	}

	// MARK: System Prompt

	systemPrompt := fmt.Sprintf(`
		YOUR ROLE:
			- You are a ai, that generate text based on the topic, with a certain length.
				- The topic is: %s.
				- The required length is %d words.
				- The response should only be valid %s, with no explanation text and %s only.
				- If the prompt is asking for code, then don't add any comments to the code.
		
		RULES:
			- Your output must contain exactly %d words — never more, never less.
			- No explanations, no notes, no extra formatting.
			- No code markdown, no examples.
   			- Do not acknowledge instructions;
	`,
		text,
		outputLengthT,
		str,
		str,
		outputLengthT,
	)
	result, err := geminiClient.Models.GenerateContent(
		ctx, "gemini-2.5-flash", gemini.Text(systemPrompt), nil)
	if err != nil {
		log.Fatal("Error generating content: ", err)
	}

	if outputFileName != nil {
		var fileName string = *outputFileName
		helpers.CreateAndWriteFile(fileName, result.Text())
		color.Green("File generated: %s", fileName)
	} else {
		color.Green("\n\nOutput: ")
		fmt.Print(result.Text())
	}

}
