package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func ExtractFieldNameFromError(errorMessage string) (fieldErrorsReturn map[string]string, boolReturn bool) {
	fieldErrors := make(map[string]string)
	// Define a regular expression pattern to match the field name in the error message
	regexPattern := `Key: '([^']+)' Error:Field validation for '([^']+)' failed on the '([^']+)' tag`
	regex := regexp.MustCompile(regexPattern)

	boolReturn = false

	// Find all matches in the error message
	matches := regex.FindAllStringSubmatch(errorMessage, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			fieldName := match[2]
			errorMessage := match[3]

			// Combine the key and field name to form a unique identifier
			identifier := fmt.Sprintf("%s", fieldName)

			// Store the error message in the map using the identifier as the key
			fieldErrors[identifier] = fmt.Sprintf("%s is %s", identifier, errorMessage)
		}

		return fieldErrors, true
	} else {
		fmt.Printf("%# v\n", errorMessage)
		// Define a regular expression pattern to match the field name in the error message
		patternErrJsonUnMarshal := `cannot unmarshal (\S|\s)+ into Go struct field (\S+) of type (\S+)`
		// get value from regexPatternJsonUnmarshallErr
		re := regexp.MustCompile(patternErrJsonUnMarshal)

		// Find matches in the error message
		matches := re.FindStringSubmatch(errorMessage)
		fmt.Printf("%# v\n", matches)
		if len(matches) > 0 {
			fieldName := matches[2]

			// Split the field name by dots and get the last part
			parts := strings.Split(fieldName, ".")
			fieldName = parts[len(parts)-1]

			fieldType := matches[3]

			// Combine the key and field name to form a unique identifier
			fieldErrors[fieldName] = fmt.Sprintf("%s is not valid, must %s type", fieldName, fieldType)

			return fieldErrors, true
		}
	}

	return fieldErrors, boolReturn
}

func GlobalCheckingErrorBindJson(errMessage string) (message string, returnError map[string]string) {
	if errMessage == "EOF" {
		message := "Request body is empty"
		return message, map[string]string{
			"error": errMessage,
		}
	}
	returnDataErrorCheck, isExistError := ExtractFieldNameFromError(errMessage)
	fmt.Printf("%# v\n", isExistError)
	if isExistError {
		message := "Validation error"
		return message, returnDataErrorCheck
	} else {
		mapReturn := map[string]string{
			"error": errMessage,
		}
		return errMessage, mapReturn
	}
}

func PaginationHelper(ctx *gin.Context) (page int, pageSize int, err error) {
	pageQuery := ctx.DefaultQuery("page", "1")
	pageSizeQuery := ctx.DefaultQuery("page_size", "10")

	page, err = ConvertStringToInt(pageQuery)
	if err != nil {
		return 0, 0, err
	}
	pageSize, err = ConvertStringToInt(pageSizeQuery)
	if err != nil {
		return 0, 0, err
	}

	return page, pageSize, nil
}
