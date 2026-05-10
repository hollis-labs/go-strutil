// Package main demonstrates the truncation helpers: Truncate, Words, Limit.
// All operate on rune counts, so multi-byte UTF-8 input is handled correctly.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	s := "the quick brown fox jumps over the lazy dog"
	fmt.Printf("Truncate(s, 15, %q)\n  = %q\n", "...", strutil.Truncate(s, 15, "..."))
	fmt.Printf("Words(s, 4, %q)\n  = %q\n", "...", strutil.Words(s, 4, "..."))
	fmt.Printf("Limit(s, 10)\n  = %q\n", strutil.Limit(s, 10))

	multi := "café résumé naïve"
	fmt.Printf("\nLimit(%q, 5) = %q  (rune-safe)\n", multi, strutil.Limit(multi, 5))
}
