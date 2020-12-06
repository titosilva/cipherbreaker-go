package cipherbreakersections

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// CipherSection struct
// Defines a section used to
// cipher or decipher text
type CipherSection struct {
	CipherType int
	Mode       int
	InputMode  int
	FilePath   string
	Key        string
	Action     int
}

// Run method of CipherSection
func (cSection CipherSection) Run() section.Section {
	// Instructions
	inst := renderable.NewContainer()
	inst.SetFixedSize(30, 6)
	inst.SetCentralization(true, false)
	inst.SetBorder('@', '@', '@', '@', '@', '@', '@', '@')
	moves := renderable.NewText("NAVIGATION\nw: up\ns: down\nEnter: select", 0, 0)
	inst.AddItem(&moves)
	inst.SetPosition(2, 0)

	if cSection.Mode == 0 {
		// cipher selection mode
		cipherSelectionView := view.NewView(nil)
		defer cipherSelectionView.Kill()

		_, height := screen.GetSize()
		h := uint(height)

		cipherSelectionView.ViewContainer.AddItem(&inst)

		// Title
		title := renderable.NewText("Select a cipher:   ", h/3-1, 0)
		cipherSelectionView.ViewContainer.AddItem(&title)

		// Options
		opt := renderable.NewOptionsList('>')

		op1 := renderable.NewText("Caesar cipher      ", 0, 0)
		op2 := renderable.NewText("Vigenere cipher    ", 1, 0)
		op3 := renderable.NewText("Exit               ", 2, 0)

		opt.AddOption(&op1)
		opt.AddOption(&op2)
		opt.AddOption(&op3)
		opt.SetPosition(h/3+1, 0)

		cipherSelectionView.ViewContainer.AddItem(&opt)

		cipherSelectionView.ViewContainer.SetCentralization(true, false)
		cipherSelectionView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
		cipherSelectionView.Show()
		go cipherSelectionView.DynamicRender()

		interacting := true
		for interacting {
			i := opt.Interact()
			switch i {
			case screen.KeyEscape:
				i = opt.Interact()
			case screen.KeyEnter:
				selected := opt.GetSelectedOption()
				switch selected {
				case 0, 1:
					return CipherSection{
						CipherType: selected,
						Mode:       1,
					}
				case 2:
					interacting = false
					return MainSection{}
				}
			default:
				i = opt.Interact()
			}
		}
	} else if cSection.Mode == 1 {
		// Input selection mode
		inputSView := view.NewView(nil)
		defer inputSView.Kill()

		_, height := screen.GetSize()
		h := uint(height)

		inputSView.ViewContainer.AddItem(&inst)

		// Title
		title := renderable.NewText("Select a mode of input:   ", h/3-1, 0)
		inputSView.ViewContainer.AddItem(&title)

		// Options
		opt := renderable.NewOptionsList('>')

		op1 := renderable.NewText("From file      ", 0, 0)
		op2 := renderable.NewText("Type           ", 1, 0)
		op3 := renderable.NewText("Exit           ", 2, 0)

		opt.AddOption(&op1)
		opt.AddOption(&op2)
		opt.AddOption(&op3)
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
				switch selected {
				case 0, 1:
					return CipherSection{
						CipherType: cSection.CipherType,
						InputMode:  selected,
						Mode:       2,
					}
				case 2:
					interacting = false
					return MainSection{}
				}
			default:
				i = opt.Interact()
			}
		}
	} else if cSection.Mode == 2 && cSection.InputMode == 0 {
		// Input from file mode
		fileModeView := view.NewView(nil)
		defer fileModeView.Kill()

		_, height := screen.GetSize()
		h := uint(height)

		fileModeView.ViewContainer.AddItem(&inst)

		title := renderable.NewText("Type the file path", h/3-1, 0)

		cipher := renderable.NewButton("> Cipher", "Cipher")
		cipher.SetPosition(0, 0)
		decipher := renderable.NewButton("> Decipher", "Decipher")
		cipher.SetPosition(0, 12)
		buttons := renderable.NewContainer()
		buttons.AddItem(&cipher)
		buttons.AddItem(&decipher)
		buttons.SetPosition(h/3+9, 0)

		userInput := renderable.NewUserInput("Path: ")
		userInput.Wrapped(40)
		userInputScrollable := renderable.NewScrollableContainer(40, 5)
		userInputScrollable.SetPosition(h/3+1, 0)
		userInputScrollable.InternalContainer.AddItem(&userInput)

		keyInput := renderable.NewUserInput("Key: ")
		if cSection.CipherType == 0 {
			keyInput.SetMaxLength(1)
		} else if cSection.CipherType == 1 {
			// Vigenere
			keyInput.Wrapped(40)
		}
		keyInputScrollable := renderable.NewScrollableContainer(40, 1)
		keyInputScrollable.SetPosition(h/3+7, 0)
		keyInputScrollable.InternalContainer.AddItem(&keyInput)

		fileModeView.ViewContainer.AddItem(&title)
		fileModeView.ViewContainer.AddItem(&buttons)
		fileModeView.ViewContainer.AddItem(&userInputScrollable)
		fileModeView.ViewContainer.AddItem(&keyInputScrollable)

		fileModeView.ViewContainer.SetCentralization(true, false)
		fileModeView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
		fileModeView.Show()
		go fileModeView.DynamicRender()

		interacting := true

		go func(updating *bool) {
			for *updating {
				userInputScrollable.ScrollDownMax()
				keyInputScrollable.ScrollDownMax()
				time.Sleep(screen.RefreshMinDelay * 10)
			}
		}(&interacting)

		var input byte
		onConfirmation := false
		interactionRouter := 0
		for interacting {
			switch interactionRouter {
			case 0:
				input = userInput.Interact()
			case 1:
				input = keyInput.Interact()
			case 2:
				input = decipher.Interact()
			case 3:
				input = cipher.Interact()
			}

			switch input {
			case 's':
				onConfirmation = !onConfirmation
				time.Sleep(screen.RefreshMinDelay)
				interactionRouter = (interactionRouter + 1) % 4
			case 'w':
				onConfirmation = !onConfirmation
				time.Sleep(screen.RefreshMinDelay)
				interactionRouter = (interactionRouter - 1) % 4
			case screen.KeyEnter:
				key := keyInput.GetValue()
				action := 0
				if cipher.Clicked() {
					action = 0
				} else if decipher.Clicked() {
					action = 1
				}
				if cipher.Clicked() || decipher.Clicked() {
					return CipherSection{
						CipherType: cSection.CipherType,
						InputMode:  cSection.InputMode,
						Mode:       cSection.Mode + 1,
						FilePath:   userInput.GetValue(),
						Key:        key,
						Action:     action,
					}
				}
			case screen.KeyEscape:
				return MainSection{}
			}
		}
	} else if cSection.Mode == 2 && cSection.InputMode == 1 {
		// Input from typing mode
		// typingModeView := view.NewView(nil)

		// plainInput := renderable.NewUserInput("Plain Text: ")
		// key := renderable.NewUserInput("Key: ")
		// cipherInput := renderable.NewUserInput("Cipher Text: ")

	} else {
		// Unknown mode will stop execution
		return EndSection{}
	}
	return EndSection{}
}

// IsEnd method of CipherSection
// Only returns false
func (cSection CipherSection) IsEnd() bool {
	return false
}
