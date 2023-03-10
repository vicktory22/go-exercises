// Package twofer exposes a ShareWith function that returns a string.
package twofer

import "fmt"

// The ShareWith function is used to create a cool string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
