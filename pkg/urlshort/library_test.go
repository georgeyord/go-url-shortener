package urlshort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomShortUrlShouldReturnAString(t *testing.T) {
	var expected string

	got := GenerateRandomShortUrl()

	assert.IsType(t, got, expected)
}

func TestGenerateRandomShortUrlShouldReturnANotEmptyString(t *testing.T) {
	got := GenerateRandomShortUrl()

	assert.Greater(t, len(got), 0)
}

func TestThatConsequentCallsOfGenerateRandomShortUrlShouldNotReturnTheSameResult(t *testing.T) {
	firstRun := GenerateRandomShortUrl()
	secondRun := GenerateRandomShortUrl()

	assert.NotEqual(t, firstRun, secondRun)
}

func TestGenerateRandomShortUrlShouldReturnStringOfTheDefaultLength(t *testing.T) {
	got := GenerateRandomShortUrl()

	assert.Equal(t, len(got), DEFAULT_LENGTH)
}

func TestGenerateRandomShortUrlShouldReturnStringOfSpecificLength(t *testing.T) {
	var length int = 6
	got := GenerateRandomShortUrlOfLength(length)

	assert.Equal(t, len(got), length)
}
