package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordsOrderAsc(t *testing.T) {
	assert := assert.New(t)

	w1 := WordScore{
		Word:  "i",
		Score: 0.0,
	}
	w2 := WordScore{
		Word:  "heart",
		Score: 0.5,
	}
	w3 := WordScore{
		Word:  "huckabees",
		Score: 0.0,
	}

	ws := []WordScore{w1, w2, w3}

	sorted := OrderAsc(ws)

	expected := []WordScore{w3, w1, w2}

	assert.Equal(expected, sorted)

}

func TestWordsOrderDesc(t *testing.T) {
	assert := assert.New(t)

	w1 := WordScore{
		Word:  "i",
		Score: 0.0,
	}
	w2 := WordScore{
		Word:  "heart",
		Score: 0.5,
	}
	w3 := WordScore{
		Word:  "huckabees",
		Score: 0.0,
	}

	ws := []WordScore{w1, w2, w3}

	sorted := OrderDesc(ws)

	expected := []WordScore{w2, w3, w1}

	assert.Equal(expected, sorted)

}
