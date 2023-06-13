package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name := "World"

	// Check if a command-line argument is provided
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}

	greeting := GenerateGreeting(name)
	fmt.Println(greeting)
}

// GenerateGreeting generates a greeting message for the given name.
func GenerateGreeting(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", strings.TrimSpace(name))
}
