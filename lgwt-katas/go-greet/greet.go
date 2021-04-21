package greet

import (
	"fmt"
	"io"
)

// Greet says hello
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
