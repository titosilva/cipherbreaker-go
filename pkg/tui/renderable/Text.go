package renderable

import "fmt"

// Text struct
// Defines a renderable text
type Text struct {
	Object
	Text string
}

// Render method of Text
// Draws the Text
func (t Text) Render() string {
	return fmt.Sprint(t.Text)
}
