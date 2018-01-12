package word

import "fmt"

type WordScore struct {
	Word  string
	Score float32
}

// TODO! add test and comment
func (w *WordScore) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("{\"%s\":%.1f}", w.Word, w.Score)
	return []byte(str), nil
}
