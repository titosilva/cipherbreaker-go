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
		"the algorithms supported. Also this project has looots of cool stuff for computer nerds. Look at this:", 1, 0)

	dynamicContainer := renderable.NewContainer()
	dynamic := renderable.NewText("", 0, 0)

	var dynamicText string
	dynamicText = "dynamic rendering is veeeeeeeeery cool"
	dynamic.SetTextFromPointer(&dynamicText)
	// dynamic.SetText(dynamicText)

	interacting := true
	go func(text *string, updating *bool) {
		for *updating {
			textBytes := []byte(*text)
			for idx, char := range textBytes {
				if !*updating {
					break
				}

				if char == ' ' {
					continue
				}

				textBytes[idx] = char - ('z' - 'Z')
				*text = string(textBytes)
				time.Sleep(100 * time.Millisecond)
				textBytes[idx] = char
				*text = string(textBytes)
			}

			if !*updating {
				break
			} else {
				time.Sleep(3 * time.Second)
			}
		}
	}(&dynamicText, &interacting)

	dynamicContainer.AddItem(&dynamic)
	dynamicContainer.SetFixedSize(uint(width-2), 3)
	dynamicContainer.SetCentralization(true, true)
	dynamicContainer.SetPosition(3, 0)

	instructionsOptions := renderable.NewText("The Cipher/Decipher option is responsible to open the options to encode or decode using Caesar or Vigenere's Cipher ", 9, 0)
	caesarOption := renderable.NewText("Selecting the Caesar option you can choose if you want to get the encoded/decoded text from a file, giving the path and typing the key that you want."+
		"After, you can select if you want to Decipher or Cipher the text using that key", 10, 0)

	welcome.Wrapped(width - 2)
	instructionsOptions.Wrapped(width - 2)
	caesarOption.Wrapped(width - 2)

	scrollable.InternalContainer.AddItem(&welcomeHeader)
	scrollable.InternalContainer.AddItem(&welcome)
	scrollable.InternalContainer.AddItem(&instructionsOptions)
	scrollable.InternalContainer.AddItem(&caesarOption)
	scrollable.InternalContainer.AddItem(&dynamicContainer)
	// scrollable.InternalContainer.AddItem(&caesarOption2)

	insView.ViewContainer.AddItem(&scrollable)

	insView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	insView.Show()
	go insView.DynamicRender()

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
