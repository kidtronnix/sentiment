package word

import "fmt"

type WordScore struct {
	Word  string
	Score float32
}

func (w *WordScore) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("{\"%s\":%f}", w.Word, w.Score)
	return []byte(str), nil
}
