package urlshort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRadmonShortUrlShouldReturnAString(t *testing.T) {
	var expected string

	got := GenerateRadmonShortUrl()

	assert.IsType(t, got, expected)
}

func TestGenerateRadmonShortUrlShouldReturnANotEmptyString(t *testing.T) {
	got := GenerateRadmonShortUrl()

	assert.Greater(t, len(got), 0)
}

func TestThatConsequentCallsOfGenerateRadmonShortUrlShouldNotReturnTheSameResult(t *testing.T) {
	firstRun := GenerateRadmonShortUrl()
	secondRun := GenerateRadmonShortUrl()

	assert.NotEqual(t, firstRun, secondRun)
}

func TestGenerateRadmonShortUrlShouldReturnStringOfTheDefaultLength(t *testing.T) {
	got := GenerateRadmonShortUrl()

	assert.Equal(t, len(got), DEFAULT_LENGTH)
}

func TestGenerateRadmonShortUrlShouldReturnStringOfSpecificLength(t *testing.T) {
	var length int = 6
	got := GenerateRadmonShortUrlOfLength(length)

	assert.Equal(t, len(got), length)
}
