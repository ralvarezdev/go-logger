package strings

import (
	gologgerseparator "github.com/ralvarezdev/go-logger/separator"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	"strings"
)

type (
	// AddCharactersFn is a function that adds characters to a string
	AddCharactersFn func(*gologgerseparator.Content, string) string
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

// FormatString returns a formatted string
func FormatString(
	contentSeparator *gologgerseparator.Content,
	content string,
	addCharactersFn AddCharactersFn,
) string {
	// Check if the addCharactersFn is nil
	if addCharactersFn == nil {
		return ""
	}

	return addCharactersFn(contentSeparator, content)
}

// FormatStatus gets the formatted status
func FormatStatus(
	contentSeparator *gologgerseparator.Content,
	status gologgerstatus.Status,
	addCharactersFn AddCharactersFn,
) string {
	return FormatString(contentSeparator, status.String(), addCharactersFn)
}

// FormatStringArray returns a string with all the strings in the array formatted
func FormatStringArray(
	multilineSeparator *gologgerseparator.Multiline,
	stringArray *[]string,
	addCharactersFn AddCharactersFn,
) string {
	// Check if the stringArray is nil or empty, or the addCharactersFn is nil
	if stringArray == nil || len(*stringArray) == 0 || addCharactersFn == nil {
		return ""
	}

	// Check if there is only one element
	if len(*stringArray) == 1 {
		return addCharactersFn(
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

		return addCharactersFn(
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
	addCharactersFn AddCharactersFn,
) string {
	mappedErrorArray := MapErrorArrayToStringArray(errorArray)
	return FormatStringArray(
		multilineSeparator,
		mappedErrorArray,
		addCharactersFn,
	)
}
