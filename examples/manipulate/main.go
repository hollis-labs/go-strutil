// Package main demonstrates the manipulation helpers: Squish, Finish, Start,
// After, Before, Between.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	fmt.Printf("Squish(%q)\n  = %q\n", "  line1\n\tline2  ", strutil.Squish("  line1\n\tline2  "))

	fmt.Printf("\nFinish(%q, %q) = %q\n", "path/to/dir", "/", strutil.Finish("path/to/dir", "/"))
	fmt.Printf("Start(%q, %q)  = %q\n", "path/to/file", "/", strutil.Start("path/to/file", "/"))

	url := "https://example.com/api/v1/widgets"
	fmt.Printf("\nAfter(url, %q)   = %q\n", "://", strutil.After(url, "://"))
	fmt.Printf("Before(url, %q)   = %q\n", "/api", strutil.Before(url, "/api"))
	fmt.Printf("Between(%q, %q, %q) = %q\n",
		"[hello world]", "[", "]", strutil.Between("[hello world]", "[", "]"))
}
