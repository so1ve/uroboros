package parser

import (
	"regexp"

	"github.com/so1ve/uroboros/internal/options"
)

type parser struct {
	options options.Options
	input   string
	// Used to signal to callers of `readWord1` whether the word
	// contained any escape sequences. This is needed because words with
	// escape sequences must not be interpreted as keywords.
	containsEsc bool
	keywords,
	reservedWords,
	reservedWordsStrict,
	reservedWordsStrictBind *regexp.Regexp
	pos,
	lineStart,
	curLine uint64
}
