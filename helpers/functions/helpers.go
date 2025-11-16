package functions

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// MARK: Global Variables
var AllowedExtensions = []string{"txt", "json"}
var AllowedOutputLengths = []string{"short", "medium", "long"}

const (
	ShortOutputLength  = 25
	MediumOutputLength = 50
	LongOutputLength   = 100
)

// MARK: Functions

func IsExtensionAllowed(fullFileName string) bool {
	parts := strings.Split(fullFileName, ".")
	for _, ext := range AllowedExtensions {
		if ext == parts[len(parts)-1] {
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

func DoesFileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func GetOutputLength(flagValue string) (int, error) {
	switch strings.ToLower(flagValue) {
	case "short":
		return ShortOutputLength, nil
	case "medium":
		return MediumOutputLength, nil
	case "long":
		return LongOutputLength, nil
	}
	return -1, fmt.Errorf("Invalid output length provided: %s", flagValue)
}

func CreateAndWriteFile(fileName string, content string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func Spinner(done chan bool) {
	frames := []rune{'|', '/', '-', '\\'}
	i := 0

	for {
		select {
		case <-done:
			return // stop spinner
		default:
			fmt.Printf("\rLoading... %c", frames[i%len(frames)])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}
