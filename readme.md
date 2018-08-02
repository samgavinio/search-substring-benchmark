## String search benchmarks

Benchmarks for searching for a single character in a string haystack.

#### Filter:
* Length >= 5 words
* contains a question mark

Because we are only searching for a single character, the average/worst case will at best have a linear run-time (given that we only search once then we discard). Micro-optimizations can be done but they are mostly negligible for our use-case.

#### What's part of the benchmark

* Naive Search
* Naive Search from the end of the string
* Rabin-Karp
* Rabin-Karp from the end of the string
* Regex

#### How to run

Pre-requisite: There is an official way of setting a workspace for Go [How to write Go code](https://golang.org/doc/code.html)

```
cd search/
go test -bench=. -benchmem
```

#### Results

```
BenchmarkNaiveSearch-4                     	100000000	        11.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkNaiveSearchWithLength-4           	100000000	        11.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNaiveSearchReversed-4             	300000000	         4.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkNaiveSearchReversedWithLength-4   	300000000	         4.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkRabinKarp-4                       	100000000	        14.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkRabinKarpWithLength-4             	100000000	        14.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkRabinKarpReversed-4               	200000000	         7.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkRabinKarpReversedWithLength-4     	200000000	         7.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegex-4                           	  200000	      7312 ns/op	   40312 B/op	      23 allocs/op
BenchmarkRegexWithLength-4                 	  200000	     10982 ns/op	   42712 B/op	      37 allocs/op
```

The results show what you would expect for a simple single character needle.

#### Where to go from here:

Benchmarking implementations for multiple character needles would be a more interesting activity.

* Rabin-karp: string.Contains() again
* Knuth–Morris–Pratt (KMP)
* Boyer-Moore
* Aho-corasick: useful if we have many patterns we are looking for. Here’s a Go implementation open-sourced by Cloudflare: https://github.com/cloudflare/ahocorasic
* Regex - we already know this is slow, but adding-in fwiw