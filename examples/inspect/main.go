// Package main demonstrates the inspection helpers: ContainsAll, ContainsAny.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	s := "the quick brown fox"
	required := []string{"quick", "fox"}
	any := []string{"mars", "venus", "fox"}

	fmt.Printf("ContainsAll(%q, %v) = %v\n", s, required, strutil.ContainsAll(s, required))
	fmt.Printf("ContainsAny(%q, %v) = %v\n", s, any, strutil.ContainsAny(s, any))

	missing := []string{"quick", "elephant"}
	fmt.Printf("ContainsAll(%q, %v) = %v\n", s, missing, strutil.ContainsAll(s, missing))
}
