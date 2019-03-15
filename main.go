package main

import (
	"fmt"

	"github.com/cvgw/secure-random-string-generator/pkg/generator"
)

func main() {
	size := 20
	fmt.Println(generator.GenerateVariable(size, 8))
	fmt.Println(generator.GenerateFromInt(size))
}
