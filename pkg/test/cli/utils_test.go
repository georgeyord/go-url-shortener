package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureOutput(t *testing.T) {
	captured := CaptureOutput(func() {
		fmt.Print("foo")
	})

	// Check message exists
	assert.Contains(t, captured, "foo")
}

func TestProvideStdin(t *testing.T) {
	var captured string

	ProvideStdin(func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			captured = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}, "foo")

	// Check message exists
	assert.Contains(t, captured, "foo")
}

func TestParseShellArgs(t *testing.T) {
	actual := ParseShellArgs("helloworld -n foo")
	assert.Equal(t, "helloworld", actual[0])
	assert.Equal(t, "-n", actual[1])
	assert.Equal(t, "foo", actual[2])
}
