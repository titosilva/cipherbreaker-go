package caesarbreaker

import (
	"github.com/titosilva/cipherbreaker-go/pkg/cipher/algorithms/caesar"
	"github.com/titosilva/cipherbreaker-go/pkg/textanalysis/algorithms/multiwordmatcher"
)

// BreakResult struct
// Represents a possible break
type BreakResult struct {
	Key       byte
	PlainText string
}

// CaesarBreaker struct
type CaesarBreaker struct {
	WordList                 []string
	KeyChannel               chan byte
	ResultChannel            chan BreakResult
	PossiblePlainText        chan string
	BreakerThreadNumber      int
	TextAnalysisThreadNumber int
	MinNumberOfMatches       int
}

// Break method of CaesarBreaker
func (caesarBreaker CaesarBreaker) Break(cipherText string) {
	for true {
		key, keyOpen := <-caesarBreaker.KeyChannel

		if !keyOpen {
			close(caesarBreaker.ResultChannel)
			break
		}

		possiblePlainText, err := caesar.Caesar{}.Decipher(cipherText, key)

		caesarBreaker.PossiblePlainText <- possiblePlainText

		if err != nil {
			continue
		}

		var analyser = multiwordmatcher.FromChannel{
			// Buffer with length equals the number of threads consuming it
			WordChannel:   make(chan string, caesarBreaker.TextAnalysisThreadNumber),
			ThreadsNumber: caesarBreaker.TextAnalysisThreadNumber,
		}

		go func() {
			for _, word := range caesarBreaker.WordList {
				analyser.WordChannel <- word
			}
			close(analyser.WordChannel)
		}()

		matchCounter := 0
		for _, match := range analyser.Analyse(cipherText) {
			matchCounter += len(match.Indexes)
		}

		if matchCounter >= caesarBreaker.MinNumberOfMatches {
			caesarBreaker.ResultChannel <- BreakResult{Key: key, PlainText: possiblePlainText}
		}
	}
}
