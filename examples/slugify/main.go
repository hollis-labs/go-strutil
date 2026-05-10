// Package main demonstrates Slugify and SlugifyN: NFKD-based URL slug
// generation with an optional rune cap.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	titles := []string{
		"Frontend Bug",
		"Café du Monde",
		"  HELLO_world  ",
		"!!! ??? !!!",
	}
	for _, t := range titles {
		fmt.Printf("Slugify(%q) = %q\n", t, strutil.Slugify(t))
	}

	long := "the quick brown fox jumps over the lazy dog"
	fmt.Printf("\nSlugifyN(%q, 12)  = %q\n", long, strutil.SlugifyN(long, 12))
	fmt.Printf("SlugifyN(title, DefaultMaxSlugLength) caps at %d runes\n",
		strutil.DefaultMaxSlugLength)
}
