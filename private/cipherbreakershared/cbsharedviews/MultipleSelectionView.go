package cbsharedviews

import (
	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// MultipleSelectionView function
func MultipleSelectionView(titleString string, options []string) (cancelled bool, selected int) {
	// Input selection mode
	inputSView := view.NewView(nil)
	defer inputSView.Kill()

	_, height := screen.GetSize()
	h := uint(height)

	// Title
	title := renderable.NewText(titleString, h/3-1, 0)
	inputSView.ViewContainer.AddItem(&title)

	// Options
	opt := renderable.NewOptionsList('>')

	for _, option := range options {
		opText := renderable.NewText(option, 0, 0)
		opt.AddOption(&opText)
	}
	opt.SetPosition(h/3+1, 0)
	inputSView.ViewContainer.AddItem(&opt)

	inputSView.ViewContainer.SetCentralization(true, false)
	inputSView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	inputSView.Show()
	go inputSView.DynamicRender()

	interacting := true
	for interacting {
		i := opt.Interact()
		switch i {
		case screen.KeyEscape:
			i = opt.Interact()
		case screen.KeyEnter:
			selected := opt.GetSelectedOption()
			return false, selected
		default:
			i = opt.Interact()
		}
	}
	return true, 0
}
