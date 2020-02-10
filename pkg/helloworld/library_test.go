package helloworld

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHelloWorldMessageShouldCapitalizeFirstLetterOfName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"foo", "Hello Foo!"},
		{"foo bar", "Hello Foo Bar!"},
		{"", "Hello World!"},
	}

	for _, c := range cases {
		got := GetHelloWorldMessage(c.in)

		assert.Equal(t, got, c.expected, fmt.Sprintf("Name should be capitalized, expected '%s', got '%s'", c.in, c.expected))
	}
}
