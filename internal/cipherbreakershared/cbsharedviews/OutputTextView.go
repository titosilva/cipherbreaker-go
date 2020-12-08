package cbsharedviews

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// OutputTextView function
func OutputTextView(text string) {
	insView := view.NewView(nil)
	defer insView.Kill()

	width, heigth := screen.GetSize()
	// w := uint(width)

	// Scrollable
	scrollable := renderable.NewScrollableContainer(width, heigth)

	// Text
	welcome := renderable.NewText(text, 1, 0)

	welcome.Wrapped(width - 2)

	scrollable.InternalContainer.AddItem(&welcome)

	insView.ViewContainer.AddItem(&scrollable)

	insView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	insView.Show()
	go insView.DynamicRender()

	interacting := true
	for interacting {
		input := scrollable.Interact()

		switch input {
		case screen.KeyEscape:
			interacting = false
		default:
			time.Sleep(screen.RefreshMinDelay)
		}
	}
}
