package cbsharedviews

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/cipher/algorithms/caesar"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// UpdatePlain function
func UpdatePlain(p *string, update *bool, pchannel chan string) {
	var ok bool
	var text string
	for {
		text, ok = <-pchannel

		if !ok {
			break
		}

		if *update {
			*p = text
		}

		time.Sleep(screen.RefreshMinDelay * 100)
	}
}

// BreakingView function
// Used to show breaking processes
func BreakingView(breaks *renderable.OptionsList, cipherTextContent, possiblePlainText *string, keys *[]byte) byte {
	breakingView := view.NewView(nil)
	defer breakingView.Kill()

	width, height := screen.GetSize()
	w := uint(width)
	h := uint(height)

	// breaks list
	breaksContainer := renderable.NewContainer()
	breaksContainer.SetFixedSize(20, h-5)
	breaksContainer.SetPosition(0, 0)

	breaksTitle := renderable.NewText("Breaks:", 0, 0)
	breaksContainer.AddItem(&breaksTitle)

	breaks.SetPosition(1, 0)
	breaksContainer.AddItem(breaks)

	// CipherTextContainer
	cipherTextContainer := renderable.NewContainer()
	cipherTextTitle := renderable.NewText("Cipher Text:", 0, 0)
	cipherText := renderable.NewText("", 1, 0)
	cipherText.SetTextFromPointer(cipherTextContent)
	cipherText.Wrapped(width - 23)

	cipherTextContainer.AddItem(&cipherTextTitle)
	cipherTextContainer.AddItem(&cipherText)

	cipherTextContainer.SetFixedSize(w-21, h/3-2)

	// PlainTextContainer
	plainTextContainer := renderable.NewContainer()
	plainTextTitle := renderable.NewText("Plain Text:", 0, 0)
	plainText := renderable.NewText("", 1, 0)
	plainText.SetTextFromPointer(possiblePlainText)
	plainText.Wrapped(width - 23)

	plainTextContainer.AddItem(&plainTextTitle)
	plainTextContainer.AddItem(&plainText)
	plainTextContainer.SetFixedSize(w-21, h/3-2)
	plainTextContainer.SetPosition(2*h/3-2, 21)

	// Add Containers to view
	breakingView.ViewContainer.AddItem(&breaksContainer)
	breakingView.ViewContainer.AddItem(&cipherTextContainer)
	breakingView.ViewContainer.AddItem(&plainTextContainer)

	breakingView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	breakingView.Show()
	go breakingView.DynamicRender()

	interacting := true
	interactionRouter := 0
	var input byte
	for interacting {
		switch interactionRouter {
		case 0:
			input = breaks.Interact()
		}

		switch input {
		case screen.KeyEnter:
			selected := breaks.GetSelectedOption()

			input, _ := caesar.Caesar{}.Decipher(*cipherTextContent, (*keys)[selected])
			OutputTextView(input)
		}
	}

	return screen.KeyEscape
}
