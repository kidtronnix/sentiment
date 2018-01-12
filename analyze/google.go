package analyze

import (
	"context"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

	"github.com/smaxwellstewart/sentiment/word"
)

type GoogleAnalyzer struct {
	Client *language.Client
	Ctx    context.Context
}

func (a *GoogleAnalyzer) Analyze(ws []string) ([]word.WordScore, error) {
	scores := []word.WordScore{}
	for _, text := range ws {
		resp, err := a.Client.AnalyzeSentiment(a.Ctx, &languagepb.AnalyzeSentimentRequest{
			Document: &languagepb.Document{
				Source: &languagepb.Document_Content{
					Content: text,
				},
				Type: languagepb.Document_PLAIN_TEXT,
			},
		})
		if err != nil {
			return nil, err
		}

		scores = append(scores, word.WordScore{Word: text, Score: resp.GetDocumentSentiment().GetScore()})
	}

	return scores, nil
}
