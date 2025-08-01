// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// Who are you going to share the cookie with
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
