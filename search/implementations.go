package search

import (
	"fmt"
	"regexp"
	"strings"
)

const MIN_LENGTH = 5
const NEEDLE = "?"
const NEEDLE_RUNE = '?'

// Naive search
func IndexRune(haystack string) bool {
	return strings.IndexRune(haystack, NEEDLE_RUNE) > 0
}

// Naive search, checks for length
func IndexRuneWithLength(haystack string) bool {
	if len(haystack) < MIN_LENGTH {
		return false
	}

	return strings.IndexRune(haystack, NEEDLE_RUNE) > 0
}

// Naive search, from the end of the string
func IndexByteReverse(haystack string) bool {
	return strings.LastIndexByte(haystack, byte(NEEDLE_RUNE)) > 0
}

// Naive search, from the end of the string. Checks for length
func IndexByteReverseWithLength(haystack string) bool {
	if len(haystack) < MIN_LENGTH {
		return false
	}

	return strings.LastIndexByte(haystack, byte(NEEDLE_RUNE)) > 0
}

// Rabin-karp
func Contains(haystack string) bool {
	return strings.Contains(haystack, NEEDLE)
}

// Rabin-karp, checks for length
func ContainsWithLength(haystack string) bool {
	if len(haystack) < MIN_LENGTH {
		return false
	}

	return strings.Contains(haystack, NEEDLE)
}

// Rabin-karp from the end of the string
func LastIndex(haystack string) bool {
	return strings.LastIndex(haystack, NEEDLE) > 0
}

// Rabin-karp from the end of the string. Checks for length
func LastIndexWithLength(haystack string) bool {
	if len(haystack) < MIN_LENGTH {
		return false
	}

	return strings.LastIndex(haystack, NEEDLE) > 0
}

// Regex
func Regex(haystack string) bool {
	pattern := fmt.Sprintf("\\%s", NEEDLE)
	r, _ := regexp.Compile(pattern)

	return r.MatchString(haystack)
}

// Regex, checks for length
func RegexWithLength(haystack string) bool {
	pattern := fmt.Sprintf("(.{%v,})(\\%s)", MIN_LENGTH-1, NEEDLE)
	r, _ := regexp.Compile(pattern)

	return r.MatchString(haystack)
}
