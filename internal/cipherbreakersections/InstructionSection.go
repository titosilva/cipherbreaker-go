package cipherbreakersections

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// InstructionSection struct
type InstructionSection struct{}

// Run method of InstructionSection
func (s InstructionSection) Run() section.Section {
	insView := view.NewView(nil)
	defer insView.Kill()

	width, _ := screen.GetSize()
	// w := uint(width)
	// Welcome
	welcomeHeader := renderable.NewText("WELCOME TO CIPHERBREAKER!", 0, 0)
	welcome := renderable.NewText("Hello! Glad you're here. With this software you may use or break some kinds of Cipher. Currently, Caesar and Vigenere cipher are "+
		"the algorithms supported.", 1, 0)

	welcome.Wrapped(width - 2)

	insView.ViewContainer.AddItem(&welcome)
	insView.ViewContainer.AddItem(&welcomeHeader)

	insView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	insView.Show()
	go insView.DynamicRender()

	interacting := true
	for interacting {
		select {
		case input := <-screen.InputChannel:
			switch input {
			case screen.KeyEscape:
				interacting = false
			default:
				time.Sleep(screen.RefreshMinDelay)
			}
		}
	}

	return MainSection{}
}

// IsEnd method
func (s InstructionSection) IsEnd() bool {
	return false
}
