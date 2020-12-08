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

		textfrompointer struct {
			active  bool
			pointer *string
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
		if t.options.textfrompointer.active {
			return *(t.options.textfrompointer.pointer)
		}

		return t.Text
	}

	if t.options.textfrompointer.active {
		content := *(t.options.textfrompointer.pointer)
		idx := 0
		lines := make([]string, 0)
		width := t.options.wrapped.width
		for (idx+1)*width < len(content) {
			lines = append(lines, content[idx*width:(idx+1)*width]+"\r")
			idx++
		}
		lines = append(lines, content[idx*width:len(content)])
		return strings.Join(lines, "\n")
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

// SetTextFromPointer method of Text
// Sets the Text content as the value in a pointer at
// render time
func (t *Text) SetTextFromPointer(pointer *string) {
	t.options.textfrompointer.active = true
	t.options.textfrompointer.pointer = pointer
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
	var content string
	if t.options.textfrompointer.active {
		content = *(t.options.textfrompointer.pointer)
	} else {
		content = t.Text
	}

	var current string
	for !t.killed {
		if t.options.textfrompointer.active {
			current = *(t.options.textfrompointer.pointer)
		} else {
			current = t.Text
		}

		if current != content && !t.killed {
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
