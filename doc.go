// Package strutil provides a small, focused collection of string manipulation
// helpers for Go, modeled loosely on Laravel's Illuminate\Support\Str facade
// but translated to idiomatic, rune-safe Go.
//
// All functions are top-level, pure, and handle empty inputs gracefully —
// there is no package-level state, no init-time side effects, and no global
// configuration. The only non-stdlib dependency is golang.org/x/text, used
// solely for Unicode normalization inside Slugify.
//
// # Function groups
//
//   - Slug and normalization: Slugify, SlugifyN, DefaultMaxSlugLength.
//   - Case conversion: SnakeCase, KebabCase, CamelCase, StudlyCase, Title,
//     UcFirst, LcFirst.
//   - Truncation: Truncate, Words, Limit.
//   - Manipulation: Squish, Finish, Start, After, Before, Between.
//   - Inspection: ContainsAll, ContainsAny.
//   - Random: Random (crypto/rand-backed alphanumeric).
//
// # Rune safety
//
// Truncation and case conversion iterate runes rather than bytes, so
// multi-byte UTF-8 strings are handled correctly. Slugify performs NFKD
// decomposition and strips combining marks, so accented characters
// transliterate to their ASCII bases (for example, "Café" → "cafe").
//
// # Slugify vs KebabCase
//
// Slugify is destructive: it lowercases, transliterates accents, and strips
// any non-ASCII alphanumeric characters. KebabCase is non-destructive: it
// preserves multi-byte characters and only changes case and separators.
//
// See the examples directory for runnable demonstrations of each group.
package strutil
