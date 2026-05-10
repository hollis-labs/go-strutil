// Package main demonstrates the case-conversion helpers: SnakeCase,
// KebabCase, CamelCase, StudlyCase, Title, UcFirst, LcFirst.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	in := "XMLParserConfig"
	fmt.Printf("input: %q\n", in)
	fmt.Printf("  SnakeCase  = %q\n", strutil.SnakeCase(in))
	fmt.Printf("  KebabCase  = %q\n", strutil.KebabCase(in))
	fmt.Printf("  CamelCase  = %q\n", strutil.CamelCase(in))
	fmt.Printf("  StudlyCase = %q\n", strutil.StudlyCase(in))

	fmt.Printf("\nTitle(%q)   = %q\n", "hello world", strutil.Title("hello world"))
	fmt.Printf("UcFirst(%q) = %q\n", "hello", strutil.UcFirst("hello"))
	fmt.Printf("LcFirst(%q) = %q\n", "Hello", strutil.LcFirst("Hello"))
}
