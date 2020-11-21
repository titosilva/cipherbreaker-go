package wordmatcher

import "strings"

// WordMatcher struct
// Declares a analyser that takes a word, and tries to match this word
// against the given text
type WordMatcher struct {
	// Word to be matched
	Word string
}

// Analyse method of WordMatcher
// Goes through all the text, trying to match word
// when complete returns a bool indicating if there was a match
// and returns where did the match occurr (-1 if no match is found)
func (wordMatcher WordMatcher) Analyse(originalText string) (bool, int) {
	// Initialize counter as 0
	counter := 0

	// Case insensitive
	text := strings.ToLower(originalText)
	lowerWord := strings.ToLower(wordMatcher.Word)

	// Iterates over all the text, and for each char, tries to
	// find a match with the given word (or any other given string)
	for idx, c := range text {
		// Uses the counter to know where it is in the word being matched
		if byte(c) == lowerWord[counter] {
			//  If it reaches the end of the word, return true
			if counter == len(lowerWord)-1 {
				return true, idx - counter
			}

			// if not, add 1 to counter
			counter++
		} else {
			// If it is not a match, set counter to 0
			counter = 0
		}
	}

	// If the execution reaches here, the code above did
	// not return. So, there was no match
	return false, -1
}
