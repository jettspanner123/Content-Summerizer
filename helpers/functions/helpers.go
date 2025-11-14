package functions

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// MARK: Global Variables
var allowedExtensions = []string{"txt", "json"}
var AllowedOutputLengths = []string{"short", "medium", "long"}

// MARK: Functions

func IsExtensionAllowed(ext string) bool {
	for _, item := range allowedExtensions {
		if item == ext {
			return true
		}
	}
	return false
}

func TerminalError(message string, wantErrorTemplate ...bool) {
	useTemplate := true
	template := "Error: "

	if len(wantErrorTemplate) > 0 {
		useTemplate = wantErrorTemplate[0]
	}

	if useTemplate {
		template = "Error: "
	} else {
		template = "       "
	}
	color.Red(fmt.Sprintf("%s %s", template, message))
}

func IsValidOutputLength(flagValue string) bool {
	for _, outputLength := range AllowedOutputLengths {
		if strings.ToLower(flagValue) == outputLength {
			return true
		}
	}
	return false
}
