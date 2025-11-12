package main

import "flag"

func main() {

	// MARK: Declaring All The Variables
	var inputFileName string
	var outputFileName string

	// MARK: Getting All The Flags For The Variables
	flag.StringVar(&inputFileName, "f", "", "Input File Name")
	flag.StringVar(&outputFileName, "o", "", "Output File Name")

	// MARK: Parsing Flags
	flag.Parse()

	// MARK: Fallback Logic
	
}
