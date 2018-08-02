package search_test

import (
	"testing"

	"github.com/samgavinio/search-substring-benchmark/search"
)

type searcher func(string) bool

func benchmarkSearch(fx searcher, b *testing.B) {
	haystack := "Hey, I want to request time off during Easter. How do I check if I have enough?"

	for n := 0; n < b.N; n++ {
		fx(haystack)
	}
}

// Naive search, byte search
func BenchmarkNaiveSearch(b *testing.B) {
	benchmarkSearch(search.IndexRune, b)
}
func BenchmarkNaiveSearchWithLength(b *testing.B) {
	benchmarkSearch(search.IndexRuneWithLength, b)
}

// Naive search, byte search from the end of the string
func BenchmarkNaiveSearchReversed(b *testing.B) {
	benchmarkSearch(search.IndexByteReverse, b)
}
func BenchmarkNaiveSearchReversedWithLength(b *testing.B) {
	benchmarkSearch(search.IndexByteReverseWithLength, b)
}

// Rabin-karp
func BenchmarkRabinKarp(b *testing.B) {
	benchmarkSearch(search.Contains, b)
}
func BenchmarkRabinKarpWithLength(b *testing.B) {
	benchmarkSearch(search.ContainsWithLength, b)
}

// Rabin-karp from the end of the string
func BenchmarkRabinKarpReversed(b *testing.B) {
	benchmarkSearch(search.LastIndex, b)
}
func BenchmarkRabinKarpReversedWithLength(b *testing.B) {
	benchmarkSearch(search.LastIndexWithLength, b)
}

// Regex
func BenchmarkRegex(b *testing.B) {
	benchmarkSearch(search.Regex, b)
}
func BenchmarkRegexWithLength(b *testing.B) {
	benchmarkSearch(search.RegexWithLength, b)
}
