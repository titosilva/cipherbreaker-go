package renderable

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// Button struct
type Button struct {
	Object
	clicked bool
	text    struct {
		selected string
		idle     string
	}
	interacting bool
	killed      bool
}

// NewButton function
// Constructor for Button
func NewButton(selected, idle string) Button {
	return Button{
		clicked: false,
		text: struct {
			selected string
			idle     string
		}{
			selected: selected,
			idle:     idle,
		},
		interacting: false,
		killed:      false,
	}
}

// Clicked method of Button
func (b *Button) Clicked() bool {
	return b.clicked
}

// SetText method of Button
func (b *Button) SetText(selected, idle string) {
	b.text.idle = idle
	b.text.selected = selected
}

// Render method of Button
func (b *Button) Render() string {
	if b.interacting {
		return b.text.selected
	}
	return b.text.idle
}

// DynamicRender method of Button
func (b *Button) DynamicRender(update chan bool) {
	interacting := b.interacting
	for !b.killed {
		if b.interacting != interacting {
			update <- true
		}

		interacting = b.interacting
		time.Sleep(screen.RefreshMinDelay)
	}
}

// Kill method of Button
func (b *Button) Kill() {
	b.killed = true
}

// Interact method of Button
func (b *Button) Interact() byte {
	b.interacting = true
	var output byte
	for b.interacting {
		select {
		case input := <-screen.InputChannel:
			switch input {
			case screen.KeyEnter:
				b.clicked = true
			}

			b.interacting = false
			output = input
		default:
			time.Sleep(screen.RefreshMinDelay)
		}
	}
	return output
}
