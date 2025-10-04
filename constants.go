package go_logger

import (
	gostringsadd "github.com/ralvarezdev/go-strings/add"
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

var (
	// HeaderSeparator is the header separator
	HeaderSeparator = gostringsseparator.NewRepeatedContent(gostringsseparator.Space)

	// StatusSeparator is the status separator
	StatusSeparator = gostringsseparator.NewRepeatedContent(gostringsseparator.Space)

	// DescriptionSeparator is the description separator
	DescriptionSeparator = gostringsseparator.NewMultiline(
		gostringsseparator.Space,
		gostringsseparator.NewLine,
		1,
	)

	// MessageSeparator is the message separator
	MessageSeparator = gostringsseparator.Space

	// AddCharactersFn is the add characters function
	AddCharactersFn = gostringsadd.Brackets
)
