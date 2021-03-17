package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

var (
	maxCount = 10
	re       = regexp.MustCompile(`,|!|\?|\.|'|"|\*`)
)

func Top10(s string) (r []string) {
	if s == "" {
		return r
	}

	// prepare string
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "\n", " ")

	// get count words
	m := make(map[string]int)
	splitString := strings.Split(s, " ")
	for _, word := range splitString {
		word = strings.TrimSpace(word)
		word = re.ReplaceAllString(word, "")
		if word == "" || word == "-" {
			continue
		}
		m[word]++
	}

	// prepare to sort
	wcArray := make([]wordCount, 0, len(m))
	for k, v := range m {
		wcArray = append(wcArray, wordCount{word: k, count: v})
	}

	sort.Slice(wcArray, func(i, j int) bool {
		if wcArray[i].count == wcArray[j].count {
			return wcArray[i].word < wcArray[j].word
		}
		return wcArray[i].count > wcArray[j].count
	})

	// cut to right size
	for _, v := range wcArray {
		r = append(r, v.word)
		if len(r) == maxCount {
			break
		}
	}

	return r
}
