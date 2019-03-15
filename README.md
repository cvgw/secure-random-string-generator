Just some crypto string generation experiments in golang

Run the benchmarks if you feel like it
```
$> go test -bench=. ./...

goos: darwin
goarch: amd64
pkg: github.com/cvgw/secure-random-string-generator/pkg/generator
BenchmarkGenerateSingleByteBufferCharLength20-4   	  500000	      2357 ns/op
BenchmarkGenerate8ByteBufferCharLength20-4        	 2000000	       816 ns/op
BenchmarkGenerate16ByteBufferCharLength20-4       	 2000000	       879 ns/op
BenchmarkGenerate32ByteBufferCharLength20-4       	 2000000	       774 ns/op
BenchmarkGenerate64ByteBufferCharLength20-4       	 1000000	      1413 ns/op
BenchmarkGenerateFromUInt64CharLength20-4         	 2000000	       652 ns/op
BenchmarkGenerateFromUInt32CharLength20-4         	 2000000	       919 ns/op
PASS
ok  	github.com/cvgw/secure-random-string-generator/pkg/generator	14.828s
```
