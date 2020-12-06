package cipherbreakersections

import (
	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// MainSection struct
type MainSection struct{}

// Run method of MainSection
func (main MainSection) Run() section.Section {
	mainView := view.NewView(nil)
	defer mainView.Kill()

	// Configure mainView
	mainView.ViewContainer.SetCentralization(true, false)

	// Title
	_, height := screen.GetSize()
	h := uint(height)
	text := renderable.NewText("{ C I P H E R  B R E A K E R }", h/3-2, 0)
	text2 := renderable.NewText("Go", h/3-1, 0)
	mainView.ViewContainer.AddItem(&text)
	mainView.ViewContainer.AddItem(&text2)

	// Instructions
	inst := renderable.NewContainer()
	inst.SetFixedSize(30, 6)
	inst.SetCentralization(true, false)
	inst.SetBorder('@', '@', '@', '@', '@', '@', '@', '@')
	moves := renderable.NewText("NAVIGATION\nw: up\ns: down\nEnter: select", 0, 0)
	inst.AddItem(&moves)
	inst.SetPosition(2, 0)

	mainView.ViewContainer.AddItem(&inst)

	// Options
	opt := renderable.NewOptionsList('>')

	op1 := renderable.NewText("Cipher/Decipher text   ", 0, 0)
	op2 := renderable.NewText("Break ciphered text    ", 1, 0)
	op3 := renderable.NewText("Exit                   ", 2, 0)

	opt.AddOption(&op1)
	opt.AddOption(&op2)
	opt.AddOption(&op3)
	opt.SetPosition(h/3+1, 0)

	mainView.ViewContainer.AddItem(&opt)

	mainView.SetBorder('|', '+', '-', '+', '|', '+', '-', '+')
	mainView.Show()
	go mainView.DynamicRender()

	interacting := true
	for interacting {
		i := opt.Interact()
		switch i {
		case screen.KeyEscape:
			i = opt.Interact()
		case screen.KeyEnter:
			selected := opt.GetSelectedOption()
			switch selected {
			case 0:
				return CipherSection{}
			case 1:
				i = opt.Interact()
			case 2:
				interacting = false
			}
		default:
			i = opt.Interact()
		}
	}

	return EndSection{}
}

// IsEnd Method of MainSection
// Only returns false
func (main MainSection) IsEnd() bool {
	return false
}
