package renderable

import "fmt"

// Text struct
// Defines a renderable text
type Text struct {
	Object
	Text string
}

func NewText(text string, row, col uint) Text {
	t := Text{Text: text}
	t.SetPosition(row, col)
	return t
}

// Render method of Text
// Draws the Text
func (t Text) Render() string {
	return fmt.Sprint(t.Text)
}
