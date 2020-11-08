package multiwordmatcher

import (
	"github.com/titosilva/cipherbreaker-go/pkg/textanalysis/algorithms/wordmatcher"
)

// privMatch struct
// Will not be exported
// This is struct works as a simple
// container to give informations about matches found
type privMatch struct {
	Word string
}

// privMatchInfo struct
// Will not be exported
// This struct works as a simple
// container to the channels used in matching
type privMatchInfo struct {
	WordChannel chan string
	Matches     chan privMatch
}

// Dry struct
// Takes a simple list of words (in the form of an array)
// and tries to match
type Dry struct {
	WordList []string
	// Number of threads to run while matching
	ThreadsNumber int
}

func matchThread(matchInfo privMatchInfo, text string) {
	for word := range matchInfo.WordChannel {
		match, _ := wordmatcher.WordMatcher{Word: word}.Analyse(text)

		if match {
			matchInfo.Matches <- privMatch{Word: word}
		}
	}
}

// Analyse method of structure dry
// Takes the word list of Dry and tries to find matches in the given text
// using concurrency
func (dry Dry) Analyse(text string) {
	// Generate Channel container
	matchInfo := privMatchInfo{
		WordChannel: make(chan string),
		Matches:     make(chan privMatch),
	}

	// Generate Word Matching Threads
	for i := 0; i < dry.ThreadsNumber; i++ {
		go matchThread(matchInfo, text)
	}

	// Output countings
}
