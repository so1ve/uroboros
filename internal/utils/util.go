package utils

import (
	"regexp"
	"strings"
)

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func WordsRegexp(words string) *regexp.Regexp {
	return regexp.MustCompile(("^(?:" + strings.ReplaceAll(words, " ", "|") + ")$"))
}
