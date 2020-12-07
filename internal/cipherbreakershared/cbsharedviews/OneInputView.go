package cbsharedviews

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// OneInputView function
// Generates a View with one input
// and handles all the interaction
func OneInputView(titleString, preludeString, leftButtonString, rightButtonString string) (
	cancelled, leftButtonClicked, rightButtonClicked bool, inputValue string) {
	// Input from file mode
	oneInputView := view.NewView(nil)
	defer oneInputView.Kill()

	_, height := screen.GetSize()
	h := uint(height)

	title := renderable.NewText(titleString, h/3-1, 0)

	left := renderable.NewButton("> "+leftButtonString, leftButtonString)
	left.SetPosition(0, 0)
	right := renderable.NewButton("> "+rightButtonString, rightButtonString)
	right.SetPosition(0, uint(len(leftButtonString)+4))
	buttons := renderable.NewContainer()
	buttons.AddItem(&left)
	buttons.AddItem(&right)
	buttons.SetPosition(h/3+7, 0)

	userInput := renderable.NewUserInput(preludeString)
	userInput.Wrapped(40)
	userInputScrollable := renderable.NewScrollableContainer(40, 5)
	userInputScrollable.SetPosition(h/3+1, 0)
	userInputScrollable.InternalContainer.AddItem(&userInput)
	userInputScrollable.AutomaticScroll(true, true)

	oneInputView.ViewContainer.AddItem(&title)
	oneInputView.ViewContainer.AddItem(&buttons)
	oneInputView.ViewContainer.AddItem(&userInputScrollable)

	oneInputView.ViewContainer.SetCentralization(true, false)
	oneInputView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	oneInputView.Show()
	go oneInputView.DynamicRender()

	interacting := true

	var input byte
	interactionRouter := 0
	for interacting {
		switch interactionRouter {
		case 0:
			input = userInput.Interact()
		case 1:
			input = right.Interact()
		case 2:
			input = left.Interact()
		}

		switch input {
		case 's':
			interactionRouter = (interactionRouter + 1) % 3
			time.Sleep(screen.RefreshMinDelay)
		case 'w':
			interactionRouter = (interactionRouter - 1) % 3
			time.Sleep(screen.RefreshMinDelay)
		case screen.KeyEnter:
			return false, right.Clicked(), left.Clicked(), userInput.GetValue()
		case screen.KeyEscape:
			return true, right.Clicked(), left.Clicked(), userInput.GetValue()
		}
	}
	return true, right.Clicked(), left.Clicked(), userInput.GetValue()
}
