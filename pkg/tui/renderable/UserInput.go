package renderable

import (
	"fmt"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// UserInput struct
// Renderable object that can get a user input
type UserInput struct {
	Object
	Prelude     string
	Value       string
	editing     bool
	interacting bool
	wrapped     struct {
		active bool
		width  int
	}
	maxlen struct {
		active bool
		value  int
	}
	killed bool
	cursor string
}

// NewUserInput function
// Constructor for UserInput
func NewUserInput(prelude string) UserInput {
	return UserInput{
		Prelude:     prelude,
		Value:       "",
		killed:      false,
		cursor:      "",
		editing:     false,
		interacting: false,
	}
}

// Editing method of UserInput
// Indicates whether the UserInput is being edited by the user
func (ui *UserInput) Editing() bool {
	return ui.editing
}

// SetMaxLength method of UserInput
func (ui *UserInput) SetMaxLength(length int) {
	ui.maxlen.active = true
	ui.maxlen.value = length
}

// GetValue method of UserInput
// Returns the text the user inserted
func (ui *UserInput) GetValue() string {
	return ui.Value
}

// SetValue method of UserInput
// Sets the value
func (ui *UserInput) SetValue(newValue string) {
	ui.Value = newValue
}

// SetPrelude method of UserInput
// Sets the prelude
func (ui *UserInput) SetPrelude(prelude string) {
	ui.Prelude = prelude
}

// GetPrelude method of UserInput
// Gets the prelude
func (ui *UserInput) GetPrelude() (prelude string) {
	return ui.Prelude
}

// Wrapped method of UserInput
// Sets wrapped mode
func (ui *UserInput) Wrapped(width int) {
	ui.wrapped.active = true
	ui.wrapped.width = width
}

// Render method of UserInput
func (ui *UserInput) Render() string {
	var text Text
	if ui.editing {
		if ui.maxlen.active && len(ui.Value) == ui.maxlen.value {
			text = NewText(ui.Prelude+ui.Value+"#", 0, 0)
		} else {
			text = NewText(ui.Prelude+ui.Value+"_", 0, 0)
		}
	} else if ui.interacting {
		text = NewText(ui.Prelude+ui.Value+"?", 0, 0)
	} else {
		text = NewText(ui.Prelude+ui.Value, 0, 0)
	}

	if ui.wrapped.active {
		text.Wrapped(ui.wrapped.width)
	}

	return text.Render()
}

// DynamicRender method of UserInput
func (ui *UserInput) DynamicRender(update chan bool) {
	value := ui.Value
	prelude := ui.Prelude
	editing := ui.editing
	interacting := ui.interacting

	for !ui.killed {
		if (ui.Value != value) || (ui.Prelude != prelude) || (ui.editing != editing) || (ui.interacting != interacting) {
			update <- true
		}

		value = ui.Value
		prelude = ui.Prelude
		editing = ui.editing
		interacting = ui.interacting

		time.Sleep(screen.RefreshMinDelay)
	}
}

// Kill method of UserInput
func (ui *UserInput) Kill() {
	ui.killed = true
}

// Interact method of input
func (ui *UserInput) Interact() byte {
	ui.interacting = true
	ui.editing = false
	var output byte
	for ui.interacting {
		select {
		case input := <-screen.InputChannel:
			switch input {
			case screen.KeyEnter:
				ui.editing = !ui.editing
			case screen.KeyEscape:
				if ui.editing {
					ui.editing = false
				} else {
					ui.interacting = false
					output = input
				}
			case screen.KeyBackspace:
				if ui.editing {
					if len(ui.Value) > 0 {
						ui.Value = ui.Value[0 : len(ui.Value)-1]
					}
				}
			default:
				if ui.editing {
					if !(len(ui.Value) == ui.maxlen.value && ui.maxlen.active) {
						ui.Value += fmt.Sprintf("%c", input)
					}
				} else {
					ui.interacting = false
					output = input
				}
			}
		}
		time.Sleep(screen.RefreshMinDelay)
	}

	ui.editing = false
	return output
}
