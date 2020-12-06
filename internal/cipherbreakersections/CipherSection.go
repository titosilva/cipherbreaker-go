package cipherbreakersections

import (
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
}

// Run method of CipherSection
func (cSection CipherSection) Run() section.Section {
	// cipher selection mode
	if cSection.Mode == 0 {
		cipherSelectionView := view.NewView(nil)
		defer cipherSelectionView.Kill()

		_, height := screen.GetSize()
		h := uint(height)

		// Instructions
		inst := renderable.NewContainer()
		inst.SetFixedSize(30, 6)
		inst.SetCentralization(true, false)
		inst.SetBorder('@', '@', '@', '@', '@', '@', '@', '@')
		moves := renderable.NewText("NAVIGATION\nw: up\ns: down\nEnter: select", 0, 0)
		inst.AddItem(&moves)
		inst.SetPosition(2, 0)

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
					i = opt.Interact()
				case 2:
					interacting = false
					return MainSection{}
				}
			default:
				i = opt.Interact()
			}
		}
	} else if cSection.Mode == 1 {

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
