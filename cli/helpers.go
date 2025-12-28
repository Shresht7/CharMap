package main

import "strings"

// -------
// HELPERS
// -------

// Helper function to strip a string of all ANSI escape codes to hopefully get the raw string
// Note: this is a basic implementation and might not handle all ANSI sequences. Consider an external library.
func stripANSI(s string) string {
	var b strings.Builder
	inANSI := false
	for _, r := range s {
		if r == '\x1b' { // ESC character
			inANSI = true
		} else if inANSI && (r == 'm' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			// End of a common ANSI escape sequence (e.g., "\x1b[...m")
			inANSI = false
		} else if !inANSI {
			b.WriteRune(r)
		}
	}
	return b.String()
}
