package renderable

import (
	"fmt"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// Text struct
// Defines a renderable text
type Text struct {
	Object
	Text   string
	killed bool
}

// NewText function
// Constructor for a Text struct
func NewText(text string, row, col uint) Text {
	t := Text{Text: text}
	t.SetPosition(row, col)
	return t
}

// Render method of Text
// Draws the Text
func (t *Text) Render() string {
	return fmt.Sprint(t.Text)
}

// SetText method of Text
// Sets the Text content
func (t *Text) SetText(content string) {
	t.Text = content
}

// DynamicRender method of Text
// Keeps watching for changes in text content.
func (t *Text) DynamicRender(update chan bool) {
	content := t.Text
	t.killed = false

	for !t.killed {
		current := t.Text
		if current != content {
			update <- true
		}

		content = current
		time.Sleep(screen.RefreshMinDelay)
	}
}

// Kill method of Text
// Kills the text dynamic rendering
func (t *Text) Kill() {
	t.killed = true
}
