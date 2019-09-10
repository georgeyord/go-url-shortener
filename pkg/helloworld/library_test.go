package helloworld

import (
	"fmt"
	"testing"
)

func TestGetHelloWorldMessageShouldCapitalizeFirstLetterOfName(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo", "Hello Foo!"},
		{"foo bar", "Hello Foo Bar!"},
	}

	for _, c := range cases {
		got := GetHelloWorldMessage(c.in)

		if got != c.want {
			t.Error(fmt.Sprintf("Name should be capitalized, expected '%s', got '%s'", c.in, c.want))
		}
	}
}
