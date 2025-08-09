package services

import (
	"strings"
)

// Normalise trims, lowers, and removes all spaces in a string.
// Example: "Great Sword" → "greatsword"
func Normalise(input string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(input)), " ", "")
}

// Normalises while retaining spacing important for searching json data files and matching
func TrimAndLower(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

// Formats the passed string to fit the EN title case naming convention
func FormatTitleCase(input string) string {
	romanMap := map[string]string{
		"i":   "I",
		"ii":  "II",
		"iii": "III",
		"iv":  "IV",
		"v":   "V",
		"vi":  "VI",
	}

	trimmed := strings.TrimSpace(input)
	words := strings.Fields(trimmed)

	for i, word := range words {
		lower := strings.ToLower(word)
		if roman, ok := romanMap[lower]; ok {
			words[i] = roman
		} else {
			words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
		}
	}

	return strings.Join(words, " ")
}

// Removes all spaces and lowers the whole string, then capitalises the first letter of each word
// Part breaker becomes PartBreaker
func CollapseTitle(s string) string {
	words := strings.Fields(s)
	joined := strings.Join(words, "")
	return strings.Title(strings.ToLower(joined))
}

// CollapseToPartbreakerStyle Just monster hunter things...
// Some names would normally be two Part Breaker are in fact one Partbreaker
func CollapseToPartbreakerStyle(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	// Capitalise first word, lowercase the rest and join
	result := strings.Title(strings.ToLower(words[0]))
	for _, word := range words[1:] {
		result += strings.ToLower(word)
	}

	return result
}
