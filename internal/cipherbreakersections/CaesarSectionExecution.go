package cipherbreakersections

import (
	"fmt"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"

	"github.com/titosilva/cipherbreaker-go/internal/cipherbreakershared/cbsharedviews"
	caesarbreaker "github.com/titosilva/cipherbreaker-go/pkg/cipherbreaker/algorithms/CaesarBreaker"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
)

// CaesarSectionExecution struct
type CaesarSectionExecution struct {
	Words                    []string
	MinMatches               int
	TextAnalysisThreadNumber int
	CipherText               string
}

func updateOpt(opt *renderable.OptionsList, keys *[]byte, breakChannel chan caesarbreaker.BreakResult) {
	for {
		b, ok := <-breakChannel

		if !ok {
			break
		}

		text := renderable.NewText(fmt.Sprintf("Key: %c", b.Key), 0, 0)
		*keys = append(*keys, b.Key)
		opt.AddOption(&text)
	}
}

// Run method of BreakerSectionExecution
func (bse CaesarSectionExecution) Run() section.Section {
	keyChannel := make(chan byte)
	breakChannel := make(chan caesarbreaker.BreakResult)
	possiblePlain := make(chan string)

	breaker := caesarbreaker.CaesarBreaker{
		WordList:                 bse.Words,
		KeyChannel:               keyChannel,
		ResultChannel:            breakChannel,
		PossiblePlainText:        possiblePlain,
		TextAnalysisThreadNumber: bse.TextAnalysisThreadNumber,
		MinNumberOfMatches:       bse.MinMatches,
	}

	optList := renderable.NewOptionsList('>')

	var plain string
	var cipherText string
	var update bool
	var keys []byte

	// PlainText update
	go cbsharedviews.UpdatePlain(&plain, &update, possiblePlain)
	// Options update
	go updateOpt(&optList, &keys, breakChannel)
	// Key generation
	go caesarbreaker.GenerateKeys(keyChannel)

	cipherText = bse.CipherText
	breaker.Break(cipherText)

	interacting := true
	for interacting {
		input := cbsharedviews.BreakingView(&optList, &cipherText, &plain, &keys)

		switch input {
		case screen.KeyEscape:
			return MainSection{}
		default:
			time.Sleep(screen.RefreshMinDelay)
		}
	}

	return MainSection{}
}

// IsEnd method of CaesarSectionExecution
func (bse CaesarSectionExecution) IsEnd() bool {
	return false
}
