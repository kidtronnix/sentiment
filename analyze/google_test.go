package analyze

import (
	"context"
	"testing"

	language "cloud.google.com/go/language/apiv1"
	"github.com/stretchr/testify/assert"
)

// NOTE! This test needs internet connection and will need google access to Language api.

func TestStringToWords(t *testing.T) {
	assert := assert.New(t)

	// if you have google default application creds setup this will add auth creds to ctx
	ctx := context.Background()

	// let client use API key (little unsure if this is actually used)
	// TODO! add client logging so can debug reqs easily.
	client, err := language.NewClient(ctx)
	if err != nil {
		assert.NoError(err)
	}

	// construct http handler

	a := &GoogleAnalyzer{
		Client: client,
		Ctx:    ctx,
	}

	ws, err := a.Analyze([]string{"love", "hate"})
	assert.NoError(err)

	// Test wordscores look correct.
	// This test is relying on google sentiment to be positive and negative for
	// 'love' and 'hate' respectively bu i think this is reasonable.
	assert.Len(ws, 2)
	for _, w := range ws {
		switch w.Word {
		case "love":
			assert.True(w.Score > 0)
		case "hate":
			assert.True(w.Score < 0)
		default:
			assert.Fail("Unexpected word!")
		}
	}
}
