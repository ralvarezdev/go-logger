package strings

import (
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	"strings"
)

// Separator is the separator for a string
type Separator string

// Separator constants
const (
	SpaceSeparator   Separator = " "
	CommaSeparator   Separator = ","
	NewLineSeparator Separator = "\n"
	TabSeparator     Separator = "\t"
	NoneSeparator    Separator = ""
)

// AddCharacters adds some characters to a string
func AddCharacters(content, leftCharacters, rightCharacters string, separator Separator) string {
	if separator == NoneSeparator {
		return leftCharacters + content + rightCharacters
	}
	return leftCharacters + string(separator) + content + string(separator) + rightCharacters
}

// AddBrackets adds brackets to a string
func AddBrackets(name string, separator Separator) string {
	return AddCharacters(name, "[", "]", separator)
}

// AddCurlyBrackets adds curly brackets to a string
func AddCurlyBrackets(name string) string {
	return AddCharacters(name, "{", "}", SpaceSeparator)
}

// AddParentheses adds parentheses to a string
func AddParentheses(name string) string {
	return AddCharacters(name, "(", ")", SpaceSeparator)
}

// FormatStatus gets the formatted status
func FormatStatus(status gologgerstatus.Status, separator Separator) string {
	return AddBrackets(status.String(), separator)
}

// FormatStringArray returns a string with all the strings in the array formatted
func FormatStringArray(outerSeparator, innerSeparator Separator, stringArray *[]string) string {
	if stringArray == nil || len(*stringArray) == 0 {
		return ""
	} else if len(*stringArray) == 1 {
		return AddBrackets((*stringArray)[0], outerSeparator)
	} else {
		var formattedDetails strings.Builder
		midSeparator := string(innerSeparator) + string(outerSeparator)

		// Add formatted details
		formattedDetails.WriteString(string(innerSeparator))
		for i, str := range *stringArray {
			formattedDetails.WriteString(str)

			if i < len(*stringArray)-1 {
				formattedDetails.WriteString(midSeparator)
			}
		}

		return AddBrackets(formattedDetails.String(), outerSeparator)
	}
}

// MapErrorArrayToStringArray maps an array of errors to an array of strings
func MapErrorArrayToStringArray(errorArray *[]error) *[]string {
	if errorArray == nil || len(*errorArray) == 0 {
		return nil
	}

	// Map the errors to strings
	stringArray := make([]string, len(*errorArray))
	for i, err := range *errorArray {
		stringArray[i] = err.Error()
	}
	return &stringArray
}

// FormatErrorArray returns a string with all the errors in the array formatted
func FormatErrorArray(outerSeparator, innerSeparator Separator, errorArray *[]error) string {
	mappedErrorArray := MapErrorArrayToStringArray(errorArray)
	return FormatStringArray(outerSeparator, innerSeparator, mappedErrorArray)
}
