package word

import "sort"

// OrderAsc orders a slice of scored words by their score from lowest to highest.
// In the case of a tie ordering happens alphabetically.
func OrderAsc(a []WordScore) []WordScore {
	sort.Sort(wordsAsc(a))
	return a
}

// OrderDesc orders a slice of scored words by their score from higest to lowest.
// In the case of a tie ordering happens alphabetically.
func OrderDesc(a []WordScore) []WordScore {
	sort.Sort(wordsDesc(a))
	return a
}

// internals for sorting

type wordsAsc []WordScore

func (a wordsAsc) Len() int      { return len(a) }
func (a wordsAsc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a wordsAsc) Less(i, j int) bool {
	if a[i].Score == a[j].Score {
		return a[i].Word < a[j].Word
	}
	return a[i].Score < a[j].Score
}

type wordsDesc []WordScore

func (a wordsDesc) Len() int      { return len(a) }
func (a wordsDesc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a wordsDesc) Less(i, j int) bool {
	if a[i].Score == a[j].Score {
		return a[i].Word < a[j].Word
	}
	return a[i].Score > a[j].Score
}
