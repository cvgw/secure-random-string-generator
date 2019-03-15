Just some crypto string generation experiments in golang

Run the benchmarks if you feel like it
```
$> go test -bench=. ./...

?   	github.com/cvgw/secure-random-string-generator	[no test files]
goos: darwin
goarch: amd64
pkg: github.com/cvgw/secure-random-string-generator/pkg/generator
BenchmarkGenerateSingleByteBufferCharLength20-4   	  500000	      2396 ns/op
BenchmarkGenerate8ByteBufferCharLength20-4        	 2000000	       805 ns/op
BenchmarkGenerate16ByteBufferCharLength20-4       	 2000000	       842 ns/op
BenchmarkGenerate32ByteBufferCharLength20-4       	 2000000	       770 ns/op
BenchmarkGenerate64ByteBufferCharLength20-4       	 1000000	      1341 ns/op
BenchmarkGenerateFromUInt64CharLength20-4         	 2000000	       656 ns/op
PASS
ok  	github.com/cvgw/secure-random-string-generator/pkg/generator	14.742s
```
