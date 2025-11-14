package main

import (
	helpers "csum/helpers/functions"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

const (
	ShortOutputLength  = 50
	MediumOutputLength = 100
	LongOutputLength   = 150
)

func main() {

	// MARK: Declaring Input Variables
	var (
		inputFileName  string
		outputFileName string
		outputFileType string
		outputLength   string
	)

	// MARK: Getting All The Flags For The Variables
	flag.StringVar(&inputFileName, "input", "", "Input File Name")
	flag.StringVar(&outputFileName, "output", "output.txt", "Output File Name")
	flag.StringVar(&outputFileType, "type", "txt", "Output File Type")
	flag.StringVar(&outputLength, "length", "Short", "Output Length")

	// MARK: Defining Custom Flag Usage
	flag.Usage = func() {
		fmt.Println("Example Usage:")
		fmt.Println("\tcsum --input input.txt")
		color.Green("\tWill create a default output.txt with the summerization.\n\n")

		fmt.Println("\tcsum --input input.txt --output output.txt")
		color.Green("\tWill create a file named output.txt with the summerization.\n\n")

		fmt.Println("\tcsum --input input.txt --output output.txt --length short")
		color.Green("\tThis will create summerization with a short output length of about 50 words.")
		color.Yellow("\tOutput Length Options:")
		color.Yellow("\t    short  = 50 words")
		color.Yellow("\t    medium = 100 words")
		color.Yellow("\t    long   = 150 words")
	}
	// MARK: Parsing Flags
	flag.Parse()

	// MARK: Fallback Logic
	if inputFileName == "" {
		helpers.TerminalError("Input File Name cannot be blank. \nRun 'csum' --help for usage.")
		return
	}

	if !helpers.IsValidOutputLength(outputLength) {
		helpers.TerminalError("Invalid Output Length. \nRun 'csum' --help for usage.")
	}

	if !helpers.IsExtensionAllowed(outputFileType) {
		helpers.TerminalError(fmt.Sprintf("Error: Output file type {%s} is not allowed", outputFileType))
		return
	}

}
