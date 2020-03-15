package converters

import "strings"

// KebabCaseStringToRegular converts kebabcase string to a string with spaces
func KebabCaseStringToRegular(s string) string {
	return strings.ReplaceAll(s, "-", " ")
}
