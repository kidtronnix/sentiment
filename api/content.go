package api

import (
	"log"
	"regexp"
	"strings"
)

// RemoveSpecialCharsRegExp
const RemoveSpecialCharsRegExp = `[^\w\s]`

// global var: pre-compiled regexp
var reg *regexp.Regexp

// ContentToWords takes user input and tokenizes text into set of words.
// See below for details of steps.
func ContentToWords(s string) []string {
	s = reg.ReplaceAllString(s, "") // remove special chars
	s = strings.ToLower(s)          // convert everything to lowercase
	ws := strings.Fields(s)         // split sentence into words
	return deduplicate(ws)          // deduplicate words
}

func compileRegex() {
	var err error
	reg, err = regexp.Compile(RemoveSpecialCharsRegExp)
	if err != nil {
		log.Fatal(err)
	}
}

func deduplicate(s []string) []string {
	m := map[string]struct{}{}
	for _, v := range s {
		if _, seen := m[v]; !seen {
			s[len(m)] = v
			m[v] = struct{}{}
		}
	}

	return s[:len(m)]
}
