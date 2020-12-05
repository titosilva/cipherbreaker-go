package cipherbreakersections

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/view"
)

// MainSection struct
type MainSection struct{}

// Run method of MainSection
func (main MainSection) Run(next *section.Section) {
	mainView := view.NewView(nil)
	defer mainView.Kill()

	hour := renderable.NewText("OK", 0, 0)
	mainView.ViewContainer.AddItem(&hour)
	mainView.Show()
	go mainView.DynamicRender()

	for true {
		hour.SetText(time.Now().String())
		input, _ := screen.ReadByte()
		if input == 27 {
			break
		}
		time.Sleep(time.Second)
	}

	next = nil
}
