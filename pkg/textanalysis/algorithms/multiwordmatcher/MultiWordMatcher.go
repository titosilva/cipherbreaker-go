package multiwordmatcher

import (
	"github.com/titosilva/cipherbreaker-go/pkg/textanalysis/algorithms/wordmatcher"
)

// privMatch struct
// Will not be exported
// This is struct works as a simple
// container to give informations about matches found
type privMatch struct {
	Word  string
	Index int
}

// privMatchChannels struct
// Will not be exported
// This struct works as a simple
// container to the channels used in matching
type privMatchChannels struct {
	WordChannel chan string
	Matches     chan privMatch
}

// MatchInfo struct
// Exported struct that contains
// information about for which indexes a word was found
// in a text
type MatchInfo struct {
	Word    string
	Indexes []int
}

// Dry struct
// Takes a simple list of words (in the form of an array)
// and tries to match
type Dry struct {
	WordList []string
	// Number of threads to run while matching
	ThreadsNumber int
}

// FromChannel struct
// Takes a simple list of words (in the form of a channel)
// and tries to match
type FromChannel struct {
	WordChannel chan string
	// Number of threads to run while matching
	ThreadsNumber int
}

func matchThread(matchChannel privMatchChannels, text string) {
	for true {
		word, open := <-matchChannel.WordChannel

		if !open {
			close(matchChannel.Matches)
			break
		}

		idxBase := 0
		for true {
			match, idx := wordmatcher.WordMatcher{Word: word}.Analyse(text[idxBase:])

			if match {
				matchChannel.Matches <- privMatch{Word: word, Index: idxBase + idx}
				idxBase += len(word) + idx
			} else {
				break
			}
		}
	}
}

// Analyse method of structure FromChannel
// Takes the word channel of FromChannel and tries to find matches in the given text
// using concurrency
func (fromCh FromChannel) Analyse(text string) []MatchInfo {
	// Generate Channel container
	matchChannel := privMatchChannels{
		// Length 100 buffer
		WordChannel: make(chan string, 100),
		Matches:     make(chan privMatch, 100),
	}

	// Using goroutinesWordChannel
	// Put words in the channel
	go func() {
		for true {
			word, isOpen := <-fromCh.WordChannel

			if !isOpen {
				close(matchChannel.WordChannel)
				break
			}

			matchChannel.WordChannel <- word
		}
	}()

	// Generate Word Matching Threads
	for i := 0; i < fromCh.ThreadsNumber; i++ {
		go matchThread(matchChannel, text)
	}

	// Output formating
	// Reads from match channel
	// and appends the index to
	// match info return
	result := make([]MatchInfo, 0)
	// Reads the channel
	for match := range matchChannel.Matches {
		// Compares a received word with the words in the channel
		matched := false
		for idx, matchInfo := range result {
			if matchInfo.Word == match.Word {
				// If the word is already in the list, just append the
				// index to the list of indexes
				result[idx].Indexes = append(result[idx].Indexes, match.Index)
				matched = true
				break
			}
		}
		if matched {
			break
		}
		// If the word is not in the list yet, add it to the list
		result = append(result, MatchInfo{
			Word:    match.Word,
			Indexes: append(make([]int, 0), match.Index),
		})
	}

	return result
}

// Analyse method of structure dry
// Takes the word list of Dry and tries to find matches in the given text
// using concurrency
func (dry Dry) Analyse(text string) []MatchInfo {
	// Generate Channel container
	matchChannel := privMatchChannels{
		// Length 100 buffer
		WordChannel: make(chan string, 100),
		Matches:     make(chan privMatch, 100),
	}

	// Put words in the channel
	go func() {
		for _, word := range dry.WordList {
			matchChannel.WordChannel <- word
		}
		close(matchChannel.WordChannel)
	}()

	// Generate Word Matching Threads
	for i := 0; i < dry.ThreadsNumber; i++ {
		go matchThread(matchChannel, text)
	}

	// Output formating
	// Reads from match channel
	// and appends the index to
	// match info return
	result := make([]MatchInfo, 0)
	// Reads the channel
	for match := range matchChannel.Matches {
		// Compares a received word with the words in the channel
		matched := false
		for idx, matchInfo := range result {
			if matchInfo.Word == match.Word {
				// If the word is already in the list, just append the
				// index to the list of indexes
				result[idx].Indexes = append(result[idx].Indexes, match.Index)
				matched = true
				break
			}
		}
		if matched {
			break
		}
		// If the word is not in the list yet, add it to the list
		result = append(result, MatchInfo{
			Word:    match.Word,
			Indexes: append(make([]int, 0), match.Index),
		})
	}

	return result
}
