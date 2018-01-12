package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentToWords(t *testing.T) {
	assert := assert.New(t)

	s := "Hello, world, world!"

	ws := ContentToWords(s)

	expected := []string{"hello", "world"}

	assert.Equal(expected, ws)
}
