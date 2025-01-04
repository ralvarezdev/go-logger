package separator

import (
	"strings"
)

type (
	// Separator is the separator for a string
	Separator string

	// Content is the separator for the content
	Content struct {
		Left  Separator
		Right Separator
	}

	// Multiline is the separator for the multiline content
	Multiline struct {
		SingleLine Separator
		Line       Separator
		TabSize    int
	}
)

// Separator constants
const (
	Space   Separator = " "
	Comma   Separator = ","
	NewLine Separator = "\n"
	Tab     Separator = "\t"
)

// NewContent creates a new content separator
func NewContent(left, right Separator) *Content {
	return &Content{
		Left:  left,
		Right: right,
	}
}

// NewRepeatedContent creates a new content separator with the same separator
func NewRepeatedContent(separator Separator) *Content {
	return NewContent(separator, separator)
}

// NewMultiline creates a new multiline separator
func NewMultiline(
	singleLine, line Separator,
	tabSize int,
) *Multiline {
	return &Multiline{
		SingleLine: singleLine,
		Line:       line,
		TabSize:    tabSize,
	}
}

// Tab returns a tab separator
func (m *Multiline) Tab() Separator {
	return Separator(strings.Repeat(string(Tab), m.TabSize))
}
