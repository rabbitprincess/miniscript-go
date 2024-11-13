package utils

func SplitString(s string, isSeparator func(c rune) bool) []string {
	// Create a slice to hold the substrings
	substrs := make([]string, 0, len(s)/2)

	// Set the initial index to zero
	start := 0

	// Iterate over the characters in the string
	for i, c := range s {
		if isSeparator(c) {
			if start < i {
				substrs = append(substrs, s[start:i]) // Append substring before separator
			}
			substrs = append(substrs, string(c)) // Append the separator itself
			start = i + 1
		}
	}

	if start < len(s) {
		substrs = append(substrs, s[start:]) // Append the last remaining substring
	}
	return substrs
}
