package word

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordScoreMarshalJSON(t *testing.T) {
	assert := assert.New(t)

	w1 := WordScore{
		Word:  "heart",
		Score: 0.1,
	}
	marshalled, err := w1.MarshalJSON()
	assert.NoError(err)

	expected := `{"heart":0.1}` //

	assert.Equal(expected, fmt.Sprintf("%s", marshalled))

}
