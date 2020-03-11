package urlshort

import (
	"github.com/thanhpk/randstr"
)

const DEFAULT_LENGTH int = 5

func GenerateRandomShortUrl() string {
	return GenerateRandomShortUrlOfLength(DEFAULT_LENGTH)
}

func GenerateRandomShortUrlOfLength(length int) string {
	if length == 0 {
		length = DEFAULT_LENGTH
	}

	return randstr.String(length)
}
