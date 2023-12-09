package helper

import (
	"fmt"
	"regexp"
	"strconv"
)

func ConvertStringToInt(str string) (int, error) {
	stringConv, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return stringConv, nil
}

func MinLengthQueryParam(str string, min int) bool {
	if str != "" && len(str) < min {
		return true
	}
	return false
}

func ExtractFieldNameFromError(errorMessage string) (fieldErrorsReturn map[string]string) {
	fieldErrors := make(map[string]string)
	// Define a regular expression pattern to match the field name in the error message
	regexPattern := `Key: '([^']+)' Error:Field validation for '([^']+)' failed on the '([^']+)' tag`
	regex := regexp.MustCompile(regexPattern)

	// Find all matches in the error message
	matches := regex.FindAllStringSubmatch(errorMessage, -1)

	// Check if a match is found

	for _, match := range matches {
		fieldName := match[2]
		errorMessage := match[3]

		// Combine the key and field name to form a unique identifier
		identifier := fmt.Sprintf("%s", fieldName)

		// Store the error message in the map using the identifier as the key
		fieldErrors[identifier] = fmt.Sprintf("%s is %s", identifier, errorMessage)
	}

	return fieldErrors
}
