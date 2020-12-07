package renderable

import (
	"strings"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// Text struct
// Defines a renderable text
type Text struct {
	Object
	Text    string
	options struct {
		wrapped struct {
			active bool
			width  int
		}
	}
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
	// If text is not wrapped, just return
	// the text
	if !t.options.wrapped.active {
		return t.Text
	}

	idx := 0
	lines := make([]string, 0)
	width := t.options.wrapped.width
	for (idx+1)*width < len(t.Text) {
		lines = append(lines, t.Text[idx*width:(idx+1)*width]+"\r")
		idx++
	}
	lines = append(lines, t.Text[idx*width:len(t.Text)])

	return strings.Join(lines, "\n")
}

// SetText method of Text
// Sets the Text content
func (t *Text) SetText(content string) {
	t.Text = content
}

// Wrapped method of Text
// Sets the Text as wrapped with line of size
// width
func (t *Text) Wrapped(width int) {
	t.options.wrapped.active = true
	t.options.wrapped.width = width
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
