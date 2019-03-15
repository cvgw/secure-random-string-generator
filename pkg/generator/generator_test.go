package generator_test

import (
	"testing"

	"github.com/cvgw/secure-random-string-generator/pkg/generator"
)

func BenchmarkGenerateSingleByteBufferCharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateSingle(20)
	}
}

func BenchmarkGenerate8ByteBufferCharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 8)
	}
}

func BenchmarkGenerate16ByteBufferCharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 16)
	}
}

func BenchmarkGenerate32ByteBufferCharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 32)
	}
}

func BenchmarkGenerate64ByteBufferCharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateVariable(20, 64)
	}
}

func BenchmarkGenerateFromUInt64CharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateFromUInt64(20)
	}
}

func BenchmarkGenerateFromUInt32CharLength20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generator.GenerateFromUInt32(20)
	}
}
