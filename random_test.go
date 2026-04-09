package strutil

import (
	"regexp"
	"testing"
	"unicode/utf8"
)

var alphanumRE = regexp.MustCompile(`^[A-Za-z0-9]*$`)

func TestRandom(t *testing.T) {
	lengths := []int{0, 1, 8, 16, 64, 128}
	for _, n := range lengths {
		got := Random(n)
		if utf8.RuneCountInString(got) != n {
			t.Errorf("Random(%d) length = %d, want %d", n, utf8.RuneCountInString(got), n)
		}
		if !alphanumRE.MatchString(got) {
			t.Errorf("Random(%d) = %q contains non-alphanumeric characters", n, got)
		}
	}
}

func TestRandomNegative(t *testing.T) {
	// Negative lengths should behave like zero (no panic, empty string).
	if got := Random(-1); got != "" {
		t.Errorf("Random(-1) = %q, want \"\"", got)
	}
}

func TestRandomUniqueness(t *testing.T) {
	// Not a strict guarantee, but two 32-char random strings colliding is
	// astronomically unlikely — if they do, something's broken.
	a := Random(32)
	b := Random(32)
	if a == b {
		t.Errorf("Random(32) produced identical values twice: %q", a)
	}
}

func TestRandomDistribution(t *testing.T) {
	// Sanity check: across many calls, the output should span the charset.
	seen := make(map[rune]bool)
	for range 200 {
		for _, r := range Random(16) {
			seen[r] = true
		}
	}
	if len(seen) < 30 {
		t.Errorf("Random distribution too narrow: only %d distinct runes seen", len(seen))
	}
}
