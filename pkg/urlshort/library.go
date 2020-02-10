package urlshort

import (
	"github.com/thanhpk/randstr"
)

const DEFAULT_LENGTH int = 5

func GenerateRadmonShortUrl() string {
	return GenerateRadmonShortUrlOfLength(DEFAULT_LENGTH)
}

func GenerateRadmonShortUrlOfLength(length int) string {
	if length == 0 {
		length = DEFAULT_LENGTH
	}

	return randstr.String(length)
}
