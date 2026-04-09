package strutil

import "strings"

// ContainsAll reports whether every substring in subs is present in s. An
// empty subs slice returns true.
//
//	ContainsAll("hello world", []string{"hello", "world"}) → true
//	ContainsAll("hello world", []string{"hello", "mars"})  → false
func ContainsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// ContainsAny reports whether at least one substring in subs is present
// in s. An empty subs slice returns false.
//
//	ContainsAny("hello world", []string{"mars", "world"}) → true
//	ContainsAny("hello world", []string{"mars", "venus"}) → false
func ContainsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
