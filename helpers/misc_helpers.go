package helpers

import (
	"regexp"
	"strings"
)

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// toGoType maps SQL types to Go types
func ToGoType(sqlType string) string {
	// Regular expression to match SQL types with optional length or precision
	re := regexp.MustCompile(`([a-zA-Z]+)(\(\d+\))?`)

	// Extract base type and optional length/precision
	matches := re.FindStringSubmatch(strings.ToUpper(sqlType))
	if len(matches) < 2 {
		return "string"
	}
	baseType := matches[1]

	switch baseType {
	case "VARCHAR", "CHAR", "NVARCHAR", "NCHAR", "CLOB", "TEXT":
		return "string"
	case "INT", "INTEGER", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
		return "int"
	case "FLOAT", "DOUBLE", "REAL", "DECIMAL", "NUMERIC":
		return "float64"
	case "DATE", "DATETIME", "TIMESTAMP", "TIME", "YEAR":
		return "time.Time"
	case "BINARY", "VARBINARY", "BLOB", "LONGBLOB", "MEDIUMBLOB", "TINYBLOB":
		return "[]byte"
	case "BOOL", "BOOLEAN":
		return "bool"
	default:
		return "string"
	}
}
