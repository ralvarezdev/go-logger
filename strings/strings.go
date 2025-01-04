package strings

import (
	gologgerseparator "github.com/ralvarezdev/go-logger/separator"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	"strings"
)

// AddCharacters adds some characters to a string
func AddCharacters(
	contentSeparator *gologgerseparator.Content,
	content, leftCharacters, rightCharacters string,
) string {
	if contentSeparator == nil {
		return strings.Join(
			[]string{leftCharacters, content, rightCharacters},
			"",
		)
	}

	return strings.Join(
		[]string{
			leftCharacters,
			string(contentSeparator.Left),
			content,
			string(contentSeparator.Right),
			rightCharacters,
		}, "",
	)
}

// AddBrackets adds brackets to a string
func AddBrackets(
	contentSeparator *gologgerseparator.Content,
	content string,
) string {
	return AddCharacters(contentSeparator, content, "[", "]")
}

// AddCurlyBrackets adds curly brackets to a string
func AddCurlyBrackets(
	contentSeparator *gologgerseparator.Content,
	content string,
) string {
	return AddCharacters(contentSeparator, content, "{", "}")
}

// AddParentheses adds parentheses to a string
func AddParentheses(
	contentSeparator *gologgerseparator.Content,
	content string,
) string {
	return AddCharacters(contentSeparator, content, "(", ")")
}

// FormatStatus gets the formatted status
func FormatStatus(
	contentSeparator *gologgerseparator.Content,
	status gologgerstatus.Status,
) string {
	return AddBrackets(contentSeparator, status.String())
}

// FormatStringArray returns a string with all the strings in the array formatted
func FormatStringArray(
	multilineSeparator *gologgerseparator.Multiline,
	stringArray *[]string,
) string {
	if stringArray == nil || len(*stringArray) == 0 {
		return ""
	}

	// Check if there is only one element
	if len(*stringArray) == 1 {
		return AddBrackets(
			gologgerseparator.NewRepeatedContent(multilineSeparator.SingleLine),
			(*stringArray)[0],
		)
	} else {
		var formattedDetails strings.Builder

		// Separators
		lineSeparator := multilineSeparator.Line
		tabSeparator := multilineSeparator.Tab()
		lineAndTabSeparator := lineSeparator + tabSeparator

		// Add formatted details
		formattedDetails.WriteString(string(tabSeparator))
		for i, str := range *stringArray {
			formattedDetails.WriteString(str)

			if i < len(*stringArray)-1 {
				formattedDetails.WriteString(string(lineAndTabSeparator))
			}
		}

		return AddBrackets(
			gologgerseparator.NewRepeatedContent(lineSeparator),
			formattedDetails.String(),
		)
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
func FormatErrorArray(
	multilineSeparator *gologgerseparator.Multiline,
	errorArray *[]error,
) string {
	mappedErrorArray := MapErrorArrayToStringArray(errorArray)
	return FormatStringArray(multilineSeparator, mappedErrorArray)
}
