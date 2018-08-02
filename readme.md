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
go test -bench=.
```

#### Results

```
BenchmarkNaiveSearch-4                     	100000000	        11.5 ns/op
BenchmarkNaiveSearchWithLength-4           	100000000	        11.5 ns/op
BenchmarkNaiveSearchReversed-4             	300000000	         4.85 ns/op
BenchmarkNaiveSearchReversedWithLength-4   	300000000	         4.74 ns/op
BenchmarkRabinKarp-4                       	100000000	        14.9 ns/op
BenchmarkRabinKarpWithLength-4             	100000000	        14.1 ns/op
BenchmarkRabinKarpReversed-4               	200000000	         7.45 ns/op
BenchmarkRabinKarpReversedWithLength-4     	200000000	         7.44 ns/op
BenchmarkRegex-4                           	  200000	      8138 ns/op
BenchmarkRegexWithLength-4                 	  100000	     12649 ns/op
```

The results show what you would expect for a simple single character needle.

#### Where to go from here:

Benchmarking implementations for multiple character needles would be a more interesting activity.

* Rabin-karp: string.Contains() again
* Knuth–Morris–Pratt (KMP)
* Boyer-Moore
* Aho-corasick: useful if we have many patterns we are looking for. Here’s a Go implementation open-sourced by Cloudflare: https://github.com/cloudflare/ahocorasic
* Regex - we already know this is slow, but adding-in fwiw