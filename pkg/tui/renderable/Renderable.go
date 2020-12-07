package renderable

// Renderable Interface
// Defines an object that can be drawn by a display
// The display will render the renderable, and make the
// proper transformations before drawing
type Renderable interface {
	Render() string
	SetPosition(uint, uint)
	GetPosition() (row uint, col uint)
}

// Object struct
// Implements a SetPosition method
type Object struct {
	position Point
}

// SetPosition method of object
func (o *Object) SetPosition(row uint, col uint) {
	o.position.row = row
	o.position.col = col
}

// GetPosition method of object
func (o *Object) GetPosition() (row uint, col uint) {
	return o.position.row, o.position.col
}
