package word

import (
	"log"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func init() {
	// Make a Regex to say we only want
	var err error
	reg, err = regexp.Compile(`[^\w\s]`) // removes all special charachters
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

func StringToWords(s string) []string {
	s = reg.ReplaceAllString(s, "") // remove special chars
	s = strings.ToLower(s)          // convert everything to lowercase
	ws := strings.Fields(s)         // split sentence into words
	return deduplicate(ws)          // deduplicate words
}
