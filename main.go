package main

import (
	"flag"
	"fmt"

	textFeature "csum/features/textOperatedInputFeature"
	helpers "csum/helpers/functions"

	"github.com/fatih/color"
)

func main() {

	// MARK: Declaring Input Variables
	var (
		inputFileName  string
		outputFileName string
		outputLength   string
		textOperator   string
	)

	// MARK: Getting All The Flags For The Variables
	flag.StringVar(&inputFileName, "input", "", "Input File Name")
	flag.StringVar(&outputFileName, "output", "output.txt", "Output File Name")
	flag.StringVar(&outputLength, "length", "short", "Output Length")
	flag.StringVar(&textOperator, "text", "", "Text Operated Input")

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
		color.Yellow("        long   = 150 words")
	}
	// MARK: Parsing Flags
	flag.Parse()

	// MARK: Fallback Logic
	if inputFileName == "" && textOperator == "" {
		helpers.TerminalError("Input File Name cannot be blank. \nRun 'csum --help' for usage.")
		return
	}

	if !helpers.DoesFileExists(inputFileName) && textOperator == "" {
		helpers.TerminalError(fmt.Sprintf("Input File Name '%s' does not exist.", inputFileName))
		return
	}

	if helpers.DoesFileExists(outputFileName) && textOperator == "" {
		helpers.TerminalError(fmt.Sprintf("Ouput File Name '%s' already exist. \nPlease change the filename using the --output flag.", outputFileName))
		return
	}

	if !helpers.IsValidOutputLength(outputLength) && textOperator == "" {
		helpers.TerminalError("Invalid Output Length. \nRun 'csum --help' for usage.")
		return
	}

	if !helpers.IsExtensionAllowed(inputFileName) && textOperator == "" {
		helpers.TerminalError("Invalid Input File Type. ")
		return
	}

	if !helpers.IsExtensionAllowed(outputFileName) {
		helpers.TerminalError("Invalid Output File Type. ")
		return
	}

	if textOperator != "" && inputFileName != "" {
		helpers.TerminalError("Both text operated input and file based input can't be used together.\n")
		color.Yellow("Either use --input or --text.")
		return
	}

	// MARK: Check which feature to call.
	var isTextOperated bool = inputFileName == "" && textOperator != ""

	// MARK: Check if write to external file
	var isWriteToExternalFile bool = isTextOperated && outputFileName != ""

	if isTextOperated {
		if isWriteToExternalFile {
			var outputFile string = "output.txt"
			//fmt.Println("Write to external file detected!")
			textFeature.TextOperatedInputFeature(textOperator, outputLength, &outputFile)
		} else {
			//fmt.Println("Write to terminal file detected!")
			textFeature.TextOperatedInputFeature(textOperator, outputLength, nil)
		}
	}
}
