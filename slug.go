package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// DefaultMaxSlugLength is a sensible default cap for slug lengths. 80 runes
// fits comfortably inside VARCHAR(100) columns, keeps URLs readable, and
// stays well under practical URL-component length limits (~255 bytes on
// most systems, 2048 for whole URLs). Callers free to choose their own
// limit via SlugifyN — this constant is a starting point, not a ceiling.
const DefaultMaxSlugLength = 80

// Slugify converts any string to a URL-safe kebab-case slug.
// It lowercases, NFKD-decomposes to strip accents ("café" → "cafe"),
// replaces runs of non-alphanumeric characters with single hyphens,
// and trims leading/trailing hyphens. Returns "" for inputs that
// normalize to nothing (e.g. "!!!", "   ", "").
//
// Examples:
//
//	Slugify("Frontend Bug")  → "frontend-bug"
//	Slugify("Café du Monde") → "cafe-du-monde"
//	Slugify("HELLO_world")   → "hello-world"
//	Slugify("  !!  ")        → ""
func Slugify(s string) string {
	if s == "" {
		return ""
	}

	// NFKD decomposition splits base characters from combining marks.
	decomposed := norm.NFKD.String(s)

	var b strings.Builder
	b.Grow(len(decomposed))
	prevHyphen := false
	for _, r := range decomposed {
		// Drop combining marks (accents).
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		r = unicode.ToLower(r)
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
		if !prevHyphen && b.Len() > 0 {
			b.WriteByte('-')
			prevHyphen = true
		}
	}

	return strings.Trim(b.String(), "-")
}

// SlugifyN behaves like Slugify but caps the result at maxRunes runes. If
// truncation would leave trailing hyphens (because the cut landed in the
// middle of a word boundary), those hyphens are stripped so the slug stays
// well-formed. maxRunes <= 0 returns "".
//
// Examples:
//
//	SlugifyN("hello world foo bar", 12) → "hello-world"
//	SlugifyN("Café du Monde", 10)        → "cafe-du-mo"
//	SlugifyN("abc", 100)                 → "abc"
//	SlugifyN("hello", 0)                 → ""
//
// Use DefaultMaxSlugLength if you don't have a specific column or URL
// constraint in mind.
func SlugifyN(s string, maxRunes int) string {
	if maxRunes <= 0 {
		return ""
	}
	slug := Slugify(s)
	if utf8.RuneCountInString(slug) <= maxRunes {
		return slug
	}
	runes := []rune(slug)
	return strings.TrimRight(string(runes[:maxRunes]), "-")
}
