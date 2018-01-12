package analyze

import "github.com/smaxwellstewart/sentiment/word"

// MockAnalyzer is a dummy analyzer that can be used in place of GoogleAnalyzer this is
// useful in test or development that only concern the api aspects.
type MockAnalyzer struct {
	Err error
}

func (a *MockAnalyzer) Analyze(ws []string) ([]word.WordScore, error) {
	if a.Err != nil {
		return nil, a.Err
	}
	out := []word.WordScore{}
	for i, w := range ws {
		out = append(out, word.WordScore{
			Word:  w,
			Score: float32(i%10) / 10,
		})
	}
	return out, nil
}
