package view

import (
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// View struct
// Represents a rendered screen
type View struct {
	ViewContainer *renderable.Container
	Killed        bool
}

// NewView function
// Constructor for a View
func NewView(c *renderable.Container) View {
	if c == nil {
		newContainer := renderable.NewContainer()
		return View{ViewContainer: &newContainer}
	}

	return View{ViewContainer: c}
}

// SetBorder method of view
func (v *View) SetBorder(
	leftBorder, topLeftCorner,
	topBorder, topRightCorner,
	rightBorder, bottomRightCorner,
	bottomBorder, bottomLeftCorner byte) (borderWasSet bool) {

	width, height := screen.GetSize()

	v.ViewContainer.SetFixedSize(uint(width), uint(height))
	return v.ViewContainer.SetBorder(leftBorder, topLeftCorner,
		topBorder, topRightCorner, rightBorder, bottomRightCorner,
		bottomBorder, bottomLeftCorner)
}

// Render method of view
// Renders the view according to the size of the screen
func (v View) Render() string {
	c := renderable.NewContainer()
	c.AddItem(v.ViewContainer)

	width, height := screen.GetSize()

	c.SetFixedSize(uint(width), uint(height))

	return c.Render()
}

// Show method of View
// Shows the rendered view in the screen
func (v View) Show() {
	if !v.Killed {
		screen.Clear()
		screen.Print(v.Render())
	}
}

// DynamicRender method of View
func (v View) DynamicRender() {
	update := make(chan bool, 100)
	v.ViewContainer.DynamicRender(update)

	// Wait for update requests
	for !v.Killed {
		for request := range update {
			// Loop till buffer is empty
			stop := false
			for !stop {
				select {
				case _ = <-update:
				default:
					stop = true
				}
				time.Sleep(screen.RefreshMinDelay / 2)
			}

			if request {
				// If received, show view on the screen
				// r := strings.ReplaceAll(v.Render(), " ", "_")
				// r := strings.ReplaceAll(v.Render(), "\n", "")
				// r := strings.ReplaceAll(v.Render(), "\n", "")
				if !v.Killed {
					r := v.Render()
					screen.Print(r)
				}
				time.Sleep(screen.RefreshMinDelay)
			}
		}
	}

	close(update)
}

// Kill method of View
func (v *View) Kill() {
	v.Killed = true
	v.ViewContainer.Kill()
}
