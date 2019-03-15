package generator

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

const (
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+!"
	// We have 64 characters which we know we can represent with 6 bits as 2^6 = 64
	mask = (1 << 6) - 1
)

func GenerateFromInt(size int) string {
	var v uint64
	err := binary.Read(rand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	cache := v
	typeSize := 64
	limit := int(typeSize) / 6
	out := make([]byte, size)
	j := 0
	for i := 0; i < size; i++ {
		if j > limit-1 {
			err := binary.Read(rand.Reader, binary.BigEndian, &v)
			if err != nil {
				log.Fatal(err)
			}
			cache = v
			j = 0
		}
		index := cache & mask
		char := characters[index]
		out[i] = char
		cache >>= 6
		j++
	}

	return string(out)
}

func GenerateVariable(size int, bufSize int) string {
	b := make([]byte, bufSize)
	out := make([]byte, size)
	j := 0
	rand.Read(b)
	for i := 0; i < size; i++ {
		if j > bufSize-1 {
			rand.Read(b)
			j = 0
		}
		input := int(b[j])

		index := input & mask
		char := characters[index]
		out[i] = char
		j++
	}
	return string(out)
}

func GenerateSingle(size int) string {
	b := make([]byte, 1)
	out := make([]byte, size)
	for i := 0; i < size; i++ {
		rand.Read(b)
		index := int(b[0]) & mask
		char := characters[index]
		out[i] = char
	}
	return string(out)
}
