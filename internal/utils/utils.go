// Package utils contains all the utility functions related to tarka compiler
package utils

// IsLetter returns true if the current character is alphabet or '_'
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// IsDigit returns true if the current character is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// IsAlphaNumeric returns true if the current character is either a letter or a digit
func IsAlphaNumeric(ch byte) bool {
	return IsLetter(ch) || IsDigit(ch)
}
