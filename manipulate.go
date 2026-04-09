package strutil

import "strings"

// Squish collapses runs of whitespace (including tabs and newlines) into
// single spaces and trims leading/trailing whitespace.
//
//	Squish("  hello   world  ") → "hello world"
//	Squish("line1\n\tline2")    → "line1 line2"
func Squish(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Finish ensures s ends with suffix, appending it if missing. If s already
// ends with suffix, returns s unchanged. An empty suffix is a no-op.
//
//	Finish("path/to/dir", "/")  → "path/to/dir/"
//	Finish("path/to/dir/", "/") → "path/to/dir/"
func Finish(s, suffix string) string {
	if suffix == "" {
		return s
	}
	if strings.HasSuffix(s, suffix) {
		return s
	}
	return s + suffix
}

// Start ensures s starts with prefix, prepending it if missing. An empty
// prefix is a no-op.
//
//	Start("path/to/file", "/")  → "/path/to/file"
//	Start("/path/to/file", "/") → "/path/to/file"
func Start(s, prefix string) string {
	if prefix == "" {
		return s
	}
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}

// After returns the substring of s after the first occurrence of search.
// If search is empty or not found, returns s unchanged.
//
//	After("hello world", " ")   → "world"
//	After("one.two.three", ".") → "two.three"
//	After("hello", "z")         → "hello"
func After(s, search string) string {
	if search == "" {
		return s
	}
	_, after, found := strings.Cut(s, search)
	if !found {
		return s
	}
	return after
}

// Before returns the substring of s before the first occurrence of search.
// If search is empty or not found, returns s unchanged.
//
//	Before("hello world", " ")   → "hello"
//	Before("one.two.three", ".") → "one"
func Before(s, search string) string {
	if search == "" {
		return s
	}
	before, _, found := strings.Cut(s, search)
	if !found {
		return s
	}
	return before
}

// Between returns the substring of s between start and end markers. It uses
// the first occurrence of start and the first occurrence of end *after*
// start. If either marker is missing, returns "".
//
//	Between("[hello]", "[", "]") → "hello"
//	Between("a-b-c", "-", "-")   → "b"
//	Between("no markers", "[", "]") → ""
func Between(s, start, end string) string {
	if start == "" || end == "" {
		return ""
	}
	_, rest, found := strings.Cut(s, start)
	if !found {
		return ""
	}
	middle, _, found := strings.Cut(rest, end)
	if !found {
		return ""
	}
	return middle
}
