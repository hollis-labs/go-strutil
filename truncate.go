package strutil

import (
	"strings"
	"unicode/utf8"
)

// Truncate returns s truncated to at most length runes (not bytes), appending
// suffix if truncation actually happened. If s fits within length, it is
// returned unchanged. length is measured in runes so multi-byte UTF-8 is safe.
//
//	Truncate("hello world", 5, "...") → "hello..."
//	Truncate("hello", 10, "...")      → "hello"
//	Truncate("café", 3, "…")          → "caf…"
func Truncate(s string, length int, suffix string) string {
	if length < 0 {
		length = 0
	}
	if utf8.RuneCountInString(s) <= length {
		return s
	}
	// Slice at rune boundary.
	runes := []rune(s)
	return string(runes[:length]) + suffix
}

// Words returns s truncated to at most count words, appending suffix if
// truncation happened. Words are split on any whitespace; leading, trailing,
// and interior runs of whitespace collapse during splitting.
//
//	Words("the quick brown fox", 2, "...") → "the quick..."
//	Words("hello", 5, "...")                → "hello"
func Words(s string, count int, suffix string) string {
	if count < 0 {
		count = 0
	}
	fields := strings.Fields(s)
	if len(fields) <= count {
		return strings.Join(fields, " ")
	}
	return strings.Join(fields[:count], " ") + suffix
}

// Limit is Truncate without a suffix — a hard rune cap. Returns s unchanged
// if it fits within length.
//
//	Limit("hello world", 5) → "hello"
func Limit(s string, length int) string {
	if length < 0 {
		length = 0
	}
	if utf8.RuneCountInString(s) <= length {
		return s
	}
	runes := []rune(s)
	return string(runes[:length])
}
