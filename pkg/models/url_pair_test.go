package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWillCreateAUrlPairWithTheCorrectAttributes(t *testing.T) {
	const short = "foo"
	const long = "bar"
	got := New(short, long)

	assert.Equal(t, got.Short, short)
	assert.Equal(t, got.Long, long)
}
