package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func handleError(message string) {
	panic(errors.New(fmt.Sprintf("Error: %s", message)))
}

func getName() (name string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

func main() {
	fmt.Println("Please enter your name.")
	name := getName()
	if name == "" {
		handleError("Name is required")
	}
	fmt.Printf("Hello %s!\n", strings.Title(name))
}
