package strutil

import (
	"testing"
	"unicode/utf8"
)

func TestDefaultMaxSlugLength(t *testing.T) {
	// Documented contract: a small, positive, sensible default.
	if DefaultMaxSlugLength <= 0 {
		t.Errorf("DefaultMaxSlugLength must be positive, got %d", DefaultMaxSlugLength)
	}
	if DefaultMaxSlugLength > 255 {
		t.Errorf("DefaultMaxSlugLength unreasonably large: %d", DefaultMaxSlugLength)
	}
}

func TestSlugifyN(t *testing.T) {
	tests := []struct {
		name string
		in   string
		max  int
		want string
	}{
		{"under limit", "hello world", 20, "hello-world"},
		{"exact fit", "hello-world", 11, "hello-world"},
		{"truncation trims trailing hyphen", "hello world foo bar", 12, "hello-world"},
		{"truncation mid-word", "hello world foo", 8, "hello-wo"},
		{"empty input", "", 10, ""},
		{"zero max", "hello world", 0, ""},
		{"negative max", "hello world", -5, ""},
		{"max larger than slug", "abc", 100, "abc"},
		{"unicode collapses before truncation", "Café du Monde", 10, "cafe-du-mo"},
		{"punctuation-only still empty", "!!!", 10, ""},
		{"exact at hyphen boundary", "one-two-three", 7, "one-two"},
		{"truncation then strip multiple trailing hyphens", "a  b  c  d  e  f", 4, "a-b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SlugifyN(tt.in, tt.max)
			if got != tt.want {
				t.Errorf("SlugifyN(%q, %d) = %q, want %q", tt.in, tt.max, got, tt.want)
			}
			if utf8.RuneCountInString(got) > tt.max && tt.max > 0 {
				t.Errorf("SlugifyN(%q, %d) = %q exceeds max length", tt.in, tt.max, got)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple two words", "Hello World", "hello-world"},
		{"unicode accents", "Café du Monde", "cafe-du-monde"},
		{"mixed case with underscore", "HELLO_world", "hello-world"},
		{"punctuation only", "  !!  ", ""},
		{"empty", "", ""},
		{"single word", "frontend", "frontend"},
		{"already slug", "hello-world", "hello-world"},
		{"multiple consecutive separators", "hello   ---___ world", "hello-world"},
		{"leading/trailing whitespace", "  Hello World  ", "hello-world"},
		{"numbers preserved", "Version 2.0", "version-2-0"},
		{"transliterated umlaut", "naïve résumé", "naive-resume"},
		{"all caps", "HELLO WORLD", "hello-world"},
		{"camelCase stays together", "helloWorld", "helloworld"},
		{"emoji dropped", "hello 🌍 world", "hello-world"},
		{"only punctuation variants", "!@#$%^&*()", ""},
		{"mixed number/letter", "abc123def", "abc123def"},
		{"tab and newline", "hello\tworld\nfoo", "hello-world-foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slugify(tt.in); got != tt.want {
				t.Errorf("Slugify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
