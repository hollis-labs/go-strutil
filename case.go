package strutil

import (
	"strings"
	"unicode"
)

// splitWords tokenizes s into logical words, handling camelCase, PascalCase,
// snake_case, kebab-case, whitespace separators, digit boundaries, and
// acronyms. Acronym handling: a run of uppercase letters followed by a
// lowercase letter breaks before the last uppercase ("XMLParser" →
// ["XML", "Parser"]). A run of uppercase at the end stays intact
// ("parseXML" → ["parse", "XML"]).
func splitWords(s string) []string {
	if s == "" {
		return nil
	}
	runes := []rune(s)
	var words []string
	var cur []rune

	flush := func() {
		if len(cur) > 0 {
			words = append(words, string(cur))
			cur = cur[:0]
		}
	}

	for i, r := range runes {
		// Any non-letter, non-digit rune is a word boundary.
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			flush()
			continue
		}

		if len(cur) == 0 {
			cur = append(cur, r)
			continue
		}

		prev := cur[len(cur)-1]

		// Transition from lower or digit to upper: new word.
		if unicode.IsUpper(r) && (unicode.IsLower(prev) || unicode.IsDigit(prev)) {
			flush()
			cur = append(cur, r)
			continue
		}

		// Inside an uppercase run, if the next rune is lowercase, end the
		// current word before this upper so acronyms split off the final
		// letter: XMLParser → XML + Parser.
		if unicode.IsUpper(r) && unicode.IsUpper(prev) &&
			i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
			flush()
			cur = append(cur, r)
			continue
		}

		cur = append(cur, r)
	}
	flush()
	return words
}

// SnakeCase converts any case style to snake_case.
//
//	SnakeCase("helloWorld")  → "hello_world"
//	SnakeCase("HelloWorld")  → "hello_world"
//	SnakeCase("hello-world") → "hello_world"
//	SnakeCase("hello world") → "hello_world"
func SnakeCase(s string) string {
	words := splitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "_")
}

// KebabCase converts any case style to kebab-case. Unlike Slugify it does
// NOT perform Unicode transliteration — multi-byte characters are preserved,
// only case and separators change.
//
//	KebabCase("helloWorld")  → "hello-world"
//	KebabCase("HelloWorld")  → "hello-world"
//	KebabCase("hello_world") → "hello-world"
func KebabCase(s string) string {
	words := splitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "-")
}

// CamelCase converts any case style to camelCase.
//
//	CamelCase("hello_world") → "helloWorld"
//	CamelCase("hello-world") → "helloWorld"
//	CamelCase("Hello World") → "helloWorld"
func CamelCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(strings.ToLower(words[0]))
	for _, w := range words[1:] {
		b.WriteString(UcFirst(strings.ToLower(w)))
	}
	return b.String()
}

// StudlyCase (aka PascalCase) converts any case style to StudlyCase.
//
//	StudlyCase("hello_world") → "HelloWorld"
//	StudlyCase("hello-world") → "HelloWorld"
//	StudlyCase("hello world") → "HelloWorld"
func StudlyCase(s string) string {
	words := splitWords(s)
	var b strings.Builder
	for _, w := range words {
		b.WriteString(UcFirst(strings.ToLower(w)))
	}
	return b.String()
}

// Title converts a string to Title Case (each space-separated word capitalized).
// Does NOT change separators — only capitalization. Underscores and hyphens
// are not treated as word boundaries.
//
//	Title("hello world") → "Hello World"
//	Title("HELLO WORLD") → "Hello World"
//	Title("hello_world") → "Hello_world"
func Title(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(s))
	atWordStart := true
	for _, r := range s {
		if unicode.IsSpace(r) {
			b.WriteRune(r)
			atWordStart = true
			continue
		}
		if atWordStart {
			b.WriteRune(unicode.ToUpper(r))
			atWordStart = false
		} else {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// UcFirst capitalizes the first rune of the string and leaves the rest
// untouched. Rune-safe for multi-byte input.
//
//	UcFirst("hello") → "Hello"
//	UcFirst("")      → ""
func UcFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// LcFirst lowercases the first rune of the string and leaves the rest
// untouched. Rune-safe for multi-byte input.
//
//	LcFirst("Hello") → "hello"
//	LcFirst("")      → ""
func LcFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
