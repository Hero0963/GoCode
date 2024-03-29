package p35

import (
	"unicode/utf8"
)

func isAValidInput(name string) bool {
	if name == "" {
		return false
	}

	return true
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
