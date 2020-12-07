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

	width, heigth := screen.GetSize()


	// Scrollable
	scrollable := renderable.NewScrollableContainer(width, heigth)

	// Welcome	
	welcomeHeader := renderable.NewText("WELCOME TO CIPHERBREAKER!", 0, 0)
	welcome := renderable.NewText("Hello! Glad you're here. With this software you may use or break some kinds of Cipher. Currently, Caesar and Vigenere cipher are "+
		"the algorithms supported.", 1, 0)

		And this project has a lot of cool stuff
		dynamic rendering is veeeeeeeery cool...

	instructionsOptions := renderable.NewText("The Cipher/Decipher option is responsible to open the options to encode or decode using Caesar or Vigenere's Cipher ", 5, 0)
	caesarOption:= renderable.NewText("Selecting the Caesar option you can choose if you want to get the encoded/decoded text from a file, giving the path and typing the key that you want."+
	"After, you can select if you want to Decipher or Cipher the text using that key", 6, 0)
	
	welcome.Wrapped(width - 2)
	instructionsOptions.Wrapped(width - 2)
	caesarOption.Wrapped(width - 2)

	scrollable.InternalContainer.AddItem(&welcomeHeader)
	scrollable.InternalContainer.AddItem(&welcome)
	scrollable.InternalContainer.AddItem(&instructionsOptions)
	scrollable.InternalContainer.AddItem(&caesarOption)
	// scrollable.InternalContainer.AddItem(&caesarOption2)

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

	return MainSection{}
}

// IsEnd method
func (s InstructionSection) IsEnd() bool {
	return false
}
