package strings

import (
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	"strings"
)

type (
	// Separator is the separator for a string
	Separator string

	// ContentSeparator is the separator for the content
	ContentSeparator struct {
		LeftSeparator  Separator
		RightSeparator Separator
	}

	// MultilineSeparator is the separator for the multiline content
	MultilineSeparator struct {
		LineSeparator Separator
		TabSize       int
	}
)

// Separator constants
const (
	SpaceSeparator   Separator = " "
	CommaSeparator   Separator = ","
	NewLineSeparator Separator = "\n"
	TabSeparator     Separator = "\t"
)

// NewContentSeparator creates a new content separator
func NewContentSeparator(leftSeparator, rightSeparator Separator) *ContentSeparator {
	return &ContentSeparator{
		LeftSeparator:  leftSeparator,
		RightSeparator: rightSeparator,
	}
}

// NewRepeatedContentSeparator creates a new content separator with the same separator
func NewRepeatedContentSeparator(separator Separator) *ContentSeparator {
	return NewContentSeparator(separator, separator)
}

// NewMultilineSeparator creates a new multiline separator
func NewMultilineSeparator(lineSeparator Separator, tabSize int) *MultilineSeparator {
	return &MultilineSeparator{
		LineSeparator: lineSeparator,
		TabSize:       tabSize,
	}
}

// TabSeparator returns a tab separator
func (m *MultilineSeparator) TabSeparator() Separator {
	return Separator(strings.Repeat(string(TabSeparator), m.TabSize))
}

// AddCharacters adds some characters to a string
func AddCharacters(content, leftCharacters, rightCharacters string, contentSeparator *ContentSeparator) string {
	if contentSeparator == nil {
		return strings.Join([]string{leftCharacters, content, rightCharacters}, "")
	}

	return strings.Join(
		[]string{
			leftCharacters,
			string(contentSeparator.LeftSeparator),
			content,
			string(contentSeparator.RightSeparator),
			rightCharacters,
		}, "",
	)
}

// AddBrackets adds brackets to a string
func AddBrackets(name string, contentSeparator *ContentSeparator) string {
	return AddCharacters(name, "[", "]", contentSeparator)
}

// AddCurlyBrackets adds curly brackets to a string
func AddCurlyBrackets(name string, contentSeparator *ContentSeparator) string {
	return AddCharacters(name, "{", "}", contentSeparator)
}

// AddParentheses adds parentheses to a string
func AddParentheses(name string, contentSeparator *ContentSeparator) string {
	return AddCharacters(name, "(", ")", contentSeparator)
}

// FormatStatus gets the formatted status
func FormatStatus(status gologgerstatus.Status, contentSeparator *ContentSeparator) string {
	return AddBrackets(status.String(), contentSeparator)
}

// FormatStringArray returns a string with all the strings in the array formatted
func FormatStringArray(multilineSeparator *MultilineSeparator, stringArray *[]string) string {
	if stringArray == nil || len(*stringArray) == 0 {
		return ""
	}

	// Separators
	lineSeparator := multilineSeparator.LineSeparator
	tabSeparator := multilineSeparator.TabSeparator()
	lineAndTabSeparator := lineSeparator + tabSeparator

	// Check if there is only one element
	if len(*stringArray) == 1 {
		return AddBrackets(
			(*stringArray)[0],
			NewContentSeparator(
				lineAndTabSeparator,
				lineSeparator,
			),
		)
	} else {
		var formattedDetails strings.Builder

		// Add formatted details
		formattedDetails.WriteString(string(tabSeparator))
		for i, str := range *stringArray {
			formattedDetails.WriteString(str)

			if i < len(*stringArray)-1 {
				formattedDetails.WriteString(string(lineAndTabSeparator))
			}
		}

		return AddBrackets(formattedDetails.String(), NewRepeatedContentSeparator(lineSeparator))
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
func FormatErrorArray(multilineSeparator *MultilineSeparator, errorArray *[]error) string {
	mappedErrorArray := MapErrorArrayToStringArray(errorArray)
	return FormatStringArray(multilineSeparator, mappedErrorArray)
}
