package helpers

import "strings"

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func CapitalizeFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

// Converts a snake_case string to CamelCase.
func ToCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := range parts {
		parts[i] = CapitalizeFirstLetter(parts[i])
	}
	return strings.Join(parts, "")
}
