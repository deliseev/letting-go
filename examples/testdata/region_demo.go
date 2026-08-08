// Package testdata holds fixtures for layouts/shortcodes/region.html
// and examples/broken_test.go; not course content.
package testdata

import "fmt"

// region: greet
func greet(name string) string {
	return fmt.Sprintf("hello, %s", name)
}

// endregion: greet

func unrelated() {}
