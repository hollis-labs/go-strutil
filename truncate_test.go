package strutil

import "testing"

func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		length int
		suffix string
		want   string
	}{
		{"truncated with ellipsis", "hello world", 5, "...", "hello..."},
		{"shorter than length", "hello", 10, "...", "hello"},
		{"equal to length", "hello", 5, "...", "hello"},
		{"unicode rune count", "café", 3, "…", "caf…"},
		{"empty", "", 5, "...", ""},
		{"zero length", "hello", 0, "...", "..."},
		{"no suffix", "hello world", 5, "", "hello"},
		{"multibyte safe", "日本語テスト", 3, "...", "日本語..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.s, tt.length, tt.suffix); got != tt.want {
				t.Errorf("Truncate(%q,%d,%q) = %q, want %q", tt.s, tt.length, tt.suffix, got, tt.want)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		count  int
		suffix string
		want   string
	}{
		{"truncated at 2", "the quick brown fox", 2, "...", "the quick..."},
		{"fewer words than count", "hello", 5, "...", "hello"},
		{"exact count", "one two three", 3, "...", "one two three"},
		{"empty", "", 3, "...", ""},
		{"zero count", "hello world", 0, "...", "..."},
		{"extra whitespace", "  the   quick  fox  ", 2, "...", "the quick..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Words(tt.s, tt.count, tt.suffix); got != tt.want {
				t.Errorf("Words(%q,%d,%q) = %q, want %q", tt.s, tt.count, tt.suffix, got, tt.want)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		length int
		want   string
	}{
		{"truncated", "hello world", 5, "hello"},
		{"under limit", "hi", 5, "hi"},
		{"exact", "hello", 5, "hello"},
		{"empty", "", 5, ""},
		{"unicode", "café", 3, "caf"},
		{"zero", "hello", 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Limit(tt.s, tt.length); got != tt.want {
				t.Errorf("Limit(%q,%d) = %q, want %q", tt.s, tt.length, got, tt.want)
			}
		})
	}
}
