// Package main demonstrates Random: a cryptographically secure alphanumeric
// string generator backed by crypto/rand.
package main

import (
	"fmt"

	strutil "github.com/hollis-labs/go-strutil"
)

func main() {
	for _, n := range []int{0, 1, 8, 32} {
		fmt.Printf("Random(%2d) = %q\n", n, strutil.Random(n))
	}
}
