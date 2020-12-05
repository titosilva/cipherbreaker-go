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
	mainView.SetBorder('|', '.', '.', '.', '|', '.', '.', '.')
	mainView.Show()
	go mainView.DynamicRender()
	defer mainView.Kill()

	interacting := true
	for interacting {
		select {
		case i := <-screen.InputChannel:
			if i == 27 {
				interacting = false
			}
		default:
			hour.SetText(time.Now().String())
			time.Sleep(screen.RefreshMinDelay)
		}
	}

	next = nil
}
