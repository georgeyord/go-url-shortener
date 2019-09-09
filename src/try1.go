package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

func handleError(message string) {
	panic(errors.New(fmt.Sprintf("Error: %s", message)))
}

func getNameInput() (name string) {
	fmt.Println("Please enter your name.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

func getHelloWorldMessage(name string) string {
	if name == "" {
		handleError("Name is required")
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}

func main() {
	name := getNameInput()
	fmt.Println(aurora.Green(getHelloWorldMessage(name)))
}
