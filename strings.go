package go_utils

import (
	"strings"
)

// ToCamelCase converts a string to camel case.
func ToCamelCase(s string) string {
	var result string

	for i, v := range s {
		if i == 0 {
			result += strings.ToLower(string(v))
			continue
		}

		if string(v) == "_" || string(v) == "-" || string(v) == " " {
			continue
		}

		if s[i-1] == '_' || s[i-1] == '-' || s[i-1] == ' ' {
			result += strings.ToUpper(string(v))
			continue
		}

		result += strings.ToLower(string(v))
	}

	return result
}

// ToPascalCase converts a string to pascal case.
func ToPascalCase(s string) string {
	return strings.ToUpper(ToCamelCase(s)[0:1]) + ToCamelCase(s)[1:]
}

// ToSnakeCase converts a string to snake case.
func ToSnakeCase(s string) string {
	var result string

	for i, v := range s {
		if i == 0 {
			result += strings.ToLower(string(v))
			continue
		}

		if string(v) == "-" || string(v) == " " {
			result += "_"
			continue
		}

		if string(v) == strings.ToUpper(string(v)) {
			if string(s[i-1]) == strings.ToUpper(string(s[i-1])) {
				result += strings.ToLower(string(v))
				continue
			}

			if s[i-1] == '_' {
				result += strings.ToLower(string(v))
				continue
			} else {
				result += "_"
				result += strings.ToLower(string(v))
				continue
			}
		}

		result += strings.ToLower(string(v))
	}

	return result
}
