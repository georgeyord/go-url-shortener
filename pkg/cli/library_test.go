package cli

import (
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

func TestPrintMessageInBlue(t *testing.T) {
	captured := captureOutput(func() {
		PrintMessage("foo", aurora.Blue)
	})

	// Check message exists
	assert.Contains(t, captured, "foo")
	// Check Blue color hex exists
	assert.Contains(t, captured, "[34m")
	// Check reset color hex exists
	assert.Contains(t, captured, "[0m")
}

func TestPrintMessageInMagenta(t *testing.T) {
	captured := captureOutput(func() {
		PrintMessage("foo", aurora.Magenta)
	})

	// Check message exists
	assert.Contains(t, captured, "foo")
	// Check Blue color hex exists
	assert.Contains(t, captured, "[35m")
	// Check reset color hex exists
	assert.Contains(t, captured, "[0m")
}

func TestGetInputWithValidInput(t *testing.T) {
	var actual string
	const input = "bar"
	captured := captureOutput(func() {
		fillStdin(
			func() {
				actual = GetInput("foo")
			}, input)
	})

	// Check message exists
	assert.Contains(t, actual, input)
	// Check message exists
	assert.Contains(t, captured, "Please enter your")
	// Check label exists
	assert.Contains(t, captured, "foo")
	// Check Blue color hex exists
	assert.Contains(t, captured, "[36m")
	// Check reset color hex exists
	assert.Contains(t, captured, "[0m")
}

func TestGetInputWithEmptyInput(t *testing.T) {
	var actual string
	const input = ""
	captured := captureOutput(func() {
		fillStdin(
			func() {
				actual = GetInput("foo")
			}, input)
	})

	// Check message exists
	assert.Contains(t, actual, input)
	// Check message exists
	assert.Contains(t, captured, "Please enter your")
	// Check label exists
	assert.Contains(t, captured, "foo")
}
