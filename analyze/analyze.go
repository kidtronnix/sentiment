package analyze

import "github.com/smaxwellstewart/sentiment/word"

type Analyzer interface {
	Analyze([]string) ([]word.WordScore, error)
}

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
