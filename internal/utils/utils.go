package utils

import (
	"strings"
	"unicode"
)


func PascalToSnake(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 5) // Pre-allocate space to minimize allocations

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func SnakeToPascal(s string) string {
	var result strings.Builder
	result.Grow(len(s)) // Pre-allocate space to minimize allocations

	capitalizeNext := true

	for _, r := range s {
		if r == '_' {
			capitalizeNext = true
		} else {
			if capitalizeNext {
				result.WriteRune(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}
