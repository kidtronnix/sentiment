package analyze

import "github.com/smaxwellstewart/sentiment/word"

type Analyzer interface {
	Analyze([]string) ([]word.WordScore, error)
}
