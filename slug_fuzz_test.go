package strutil

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// checkSlugInvariants asserts the structural contract every slug must meet,
// regardless of input: ASCII-only, lowercased, no leading/trailing hyphen,
// no adjacent hyphens.
func checkSlugInvariants(t *testing.T, label, in, got string) {
	t.Helper()
	for _, r := range got {
		if !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-') {
			t.Errorf("%s(%q) = %q contains invalid rune %q", label, in, got, r)
			return
		}
	}
	if strings.HasPrefix(got, "-") {
		t.Errorf("%s(%q) = %q has leading hyphen", label, in, got)
	}
	if strings.HasSuffix(got, "-") {
		t.Errorf("%s(%q) = %q has trailing hyphen", label, in, got)
	}
	if strings.Contains(got, "--") {
		t.Errorf("%s(%q) = %q has consecutive hyphens", label, in, got)
	}
}

func FuzzSlugify(f *testing.F) {
	seeds := []string{
		"",
		"hello world",
		"Café du Monde",
		"HELLO_world",
		"!!!",
		"   ",
		"XMLParser",
		"日本語",
		"hello 🌍 world",
		"a-b-c",
		"---",
		"naïve résumé",
		"\x00\x01\x02",
		strings.Repeat("x", 500),
	}
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		got := Slugify(s)
		checkSlugInvariants(t, "Slugify", s, got)

		// Idempotence: Slugify(Slugify(x)) == Slugify(x).
		if got2 := Slugify(got); got2 != got {
			t.Errorf("Slugify not idempotent: Slugify(%q)=%q, Slugify(%q)=%q", s, got, got, got2)
		}
	})
}

func FuzzSlugifyN(f *testing.F) {
	seeds := []struct {
		s string
		n int
	}{
		{"", 10},
		{"hello world", 5},
		{"hello world foo bar", 12},
		{"Café du Monde", 10},
		{"!!!", 80},
		{"a-very-long-string-with-many-words", 8},
		{"XMLParser", 3},
		{strings.Repeat("abc def ", 50), 40},
		{"unicode 日本語 mixed", 20},
	}
	for _, seed := range seeds {
		f.Add(seed.s, seed.n)
	}
	f.Fuzz(func(t *testing.T, s string, n int) {
		got := SlugifyN(s, n)

		// n <= 0 must always yield an empty slug.
		if n <= 0 {
			if got != "" {
				t.Errorf("SlugifyN(%q, %d) = %q, want \"\" for non-positive n", s, n, got)
			}
			return
		}

		// Length cap is measured in runes.
		if count := utf8.RuneCountInString(got); count > n {
			t.Errorf("SlugifyN(%q, %d) = %q exceeds max length (%d runes)", s, n, got, count)
		}

		checkSlugInvariants(t, "SlugifyN", s, got)
	})
}
