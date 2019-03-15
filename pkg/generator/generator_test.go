package generator_test

import (
	"testing"

	"github.com/cvgw/secure-random-string-generator/pkg/generator"
)

func BenchmarkGenerateSingleByteBuffer20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateSingle(20)
	}
}

func BenchmarkGenerate8ByteBuffer20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 8)
	}
}

func BenchmarkGenerate16ByteBuffer20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 16)
	}
}

func BenchmarkGenerate32ByteBuffer20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 32)
	}
}

func BenchmarkGenerate64ByteBuffer20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 64)
	}
}

func BenchmarkGenerateFromInt20Chars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 64)
	}
}
