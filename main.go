package main

import (
	helpers "csum/helpers/functions"
	"flag"
	"fmt"
	"strings"
)

func main() {

	// MARK: Declaring Input Variables
	var inputFileName string
	var outputFileName string
	var outputFileType string

	// MARK: Getting All The Flags For The Variables
	flag.StringVar(&inputFileName, "f", "", "Input File Name")
	flag.StringVar(&outputFileName, "o", "output.txt", "Output File Name")
	flag.StringVar(&outputFileType, "ot", "txt", "Output File Type")

	// MARK: Parsing Flags
	flag.Parse()

	// MARK: Fallback Logic
	if strings.Compare(inputFileName, "") == 0 {
		helpers.TerminalError("Input File Name cannot be blank.")
		helpers.TerminalError("Run 'csum --help' for help", false)
		return
	}

	if !helpers.IsExtensionAllowed(outputFileType) {
		helpers.TerminalError(fmt.Sprintf("Error: Output file type {%s} is not allowed", outputFileType))
		return
	}

}
