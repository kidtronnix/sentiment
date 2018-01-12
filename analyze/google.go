package analyze

import (
	"context"
	"log"
	"sync"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

	"github.com/smaxwellstewart/sentiment/word"
)

const MaxParallelWorkers = 100

type GoogleAnalyzer struct {
	Client *language.Client
	Ctx    context.Context
}

func (a *GoogleAnalyzer) Analyze(ws []string) ([]word.WordScore, error) {
	scores := []word.WordScore{}

	// basic fan-out, fan-in pattern.
	// https://blog.golang.org/pipelines
	in := a.gen(ws)

	numWorkers := len(ws)
	if numWorkers > MaxParallelWorkers {
		numWorkers = MaxParallelWorkers
	}

	// start workers to score words
	cs := make([]<-chan word.WordScore, numWorkers)
	for i := 0; i < numWorkers; i++ {
		cs[i] = a.score(in)
	}

	// merge results
	for w := range a.merge(cs) {
		scores = append(scores, w)
	}

	return scores, nil
}

// send all words on output channel
func (a *GoogleAnalyzer) gen(words []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, w := range words {
			out <- w
		}
		close(out)
	}()
	return out
}

func (a *GoogleAnalyzer) score(in <-chan string) <-chan word.WordScore {
	out := make(chan word.WordScore)
	go func() {
		for w := range in {
			// google logix
			resp, err := a.Client.AnalyzeSentiment(a.Ctx, &languagepb.AnalyzeSentimentRequest{
				Document: &languagepb.Document{
					Source: &languagepb.Document_Content{
						Content: w,
					},
					Type: languagepb.Document_PLAIN_TEXT,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			out <- word.WordScore{Word: w, Score: resp.GetDocumentSentiment().GetScore()}
		}
		close(out)
	}()
	return out
}

func (a *GoogleAnalyzer) merge(cs []<-chan word.WordScore) <-chan word.WordScore {
	var wg sync.WaitGroup
	out := make(chan word.WordScore)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan word.WordScore) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// not used but left for bechnmarking
func (a *GoogleAnalyzer) naiveAnalyze(ws []string) ([]word.WordScore, error) {
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
