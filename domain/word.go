package domain

import (
	"github.com/gobeam/stringy"
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToCamelCaseFistLarge(s string) string {
	return stringy.New(s).CamelCase()
}

func ToCamelCaseFistLower(s string) string {
	rs := stringy.New(s).CamelCase()
	if len(rs) > 0 {
		return string(unicode.ToLower(rune(rs[0]))) + rs[1:]
	}
	return rs
}

func ToSnakeLower(s string) string {
	rs := stringy.New(s).SnakeCase().ToLower()
	return rs
}

func ToSnakeLarge(s string) string {
	rs := stringy.New(s).SnakeCase().ToUpper()
	return rs
}
