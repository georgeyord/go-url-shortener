package main

import (
	"fmt"
	"testing"
)

func TestGetHelloWorldMessageShouldCapitalizeFirstLetterOfName(t *testing.T) {
	capitalized := getHelloWorldMessage("foo")
	const expectedCapitalized1 = "Hello Foo!"

	if capitalized != expectedCapitalized1 {
		t.Error(fmt.Sprintf("Name should be capitalized, expected '%s', got '%s'", expectedCapitalized1, capitalized))
	}

	capitalized = getHelloWorldMessage("foo bar")
	const expectedCapitalized2 = "Hello Foo Bar!"

	if capitalized != expectedCapitalized2 {
		t.Error(fmt.Sprintf("Name should be capitalized, expected '%s', got '%s'", expectedCapitalized2, capitalized))
	}
}
