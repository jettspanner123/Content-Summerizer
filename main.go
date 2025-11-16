package main

import (
	helpers "csum/helpers/functions"
	"flag"
	"fmt"

	textOperatedInputFeature "csum/features/textOperatedInputFeature"

	"github.com/fatih/color"
)

func main() {

	// MARK: Declaring Input Variables
	var (
		inputFileName  string
		outputFileName string
		outputLength   string
		textOperator   string
		structure      string
	)

	loading := make(chan bool)

	// MARK: Getting All The Flags For The Variables
	flag.StringVar(&inputFileName, "input", "", "Input File Name")
	flag.StringVar(&outputFileName, "output", "", "Output File Name")
	flag.StringVar(&outputLength, "length", "short", "Output Length")
	flag.StringVar(&textOperator, "text", "", "Text Operated Input")
	flag.StringVar(&structure, "struct", "text", "Text Operated Input")

	// MARK: Defining Custom Flag Usage
	flag.Usage = func() {
		fmt.Println("Example Usage:")
		fmt.Println("    csum --input input.txt")
		color.Green("    Will create a default output.txt with the summerization.")
		color.Yellow("    Supported Extensions:")
		color.Yellow("        txt")
		color.Yellow("        json\n\n")

		fmt.Println("    csum --input input.txt --output output.txt")
		color.Green("    Will create a file named 'output.txt' with the summerization.")
		color.Yellow("    Supported Extensions:")
		color.Yellow("        txt")
		color.Yellow("        json\n\n")

		fmt.Println("    csum --input input.txt --output output.txt --length short")
		color.Green("    This will create summerization with a short output length of about 50 words.")
		color.Yellow("    Output Length Options:")
		color.Yellow("        short  = 50 words")
		color.Yellow("        medium = 100 words")
		color.Yellow("        long   = 150 words\n\n")

		fmt.Println("    csum --text 'Hello World in C++'")
		color.Green("    This will answer the text in the terminal itself.\n\n")

		fmt.Println("    csum --text 'Hello World in C++' --output output.txt")
		color.Green("    This will create a file named 'output.txt' with the response.")
	}
	// MARK: Parsing Flags
	flag.Parse()

	// MARK: Validation Logic

	var isTextInput bool = inputFileName == "" && textOperator != ""

	go helpers.Spinner(loading)

	if isTextInput {
		var isFileOutput bool = outputFileName != ""

		if isFileOutput {
			textOperatedInputFeature.TextOperatedInputFeature(textOperator, outputLength, &outputFileName, structure)
		} else {
			textOperatedInputFeature.TextOperatedInputFeature(textOperator, outputLength, nil, structure)
		}
	} else {

	}

	loading <- true

}
