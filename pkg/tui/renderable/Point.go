package renderable

// Point struct
// Defines an structure for coordinates
type Point struct {
	row uint
	col uint
}

// NewPoint function
// Returns a point
func NewPoint(row uint, col uint) Point {
	return Point{row: row, col: col}
}
