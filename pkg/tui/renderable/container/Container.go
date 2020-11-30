package container

import (
	"strings"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
)

// Container struct
// Container for Renderables
type Container struct {
	renderable.Object
	items   []renderable.Renderable
	options struct {
		fixedSize struct {
			active bool
			height uint
			width  uint
		}

		border struct {
			active bool
			size   int
			// Sides
			topBorderChar    byte
			bottomBorderChar byte
			leftBorderChar   byte
			rightBorderChar  byte
			// Corners
			tlCornerChar byte
			trCornerChar byte
			blCornerChar byte
			brCornerChar byte
		}
	}
}

// NewContainer function
// Returns an empty container
func NewContainer() Container {
	cont := Container{}
	return cont
}

// AddItem method of Container
// Adds a renderable item
func (c *Container) AddItem(item renderable.Renderable) {
	c.items = append(c.items, item)
}

// SetFixedSize method of Container
// It sets a fixed size to a container
func (c *Container) SetFixedSize(width uint, height uint) {
	c.options.fixedSize.active = true
	c.options.fixedSize.width = width
	c.options.fixedSize.height = height
}

// Render method of container
func (c Container) Render() (containerRendered string) {
	containerRendered = ""
	containerLines := make([]string, 0)
	for _, item := range c.items {
		var renderedLines = strings.Split(item.Render(), "\n")

		rowUint, colUint := item.GetPosition()

		row := int(rowUint)
		col := int(colUint)

		for i := row; i < row+len(renderedLines); i++ {
			// Pad the matrix so it will have, at least, the enough
			// number of lines to render the next line of the item
			// and the number of columns needed
			for len(containerLines) <= i {
				containerLines = append(containerLines, "")
			}

			for len(containerLines[i]) < col+len(renderedLines[i-row]) {
				containerLines[i] += " "
			}

			temp := []byte(containerLines[i])
			for j := col; j < col+len(renderedLines[i-row]); j++ {
				char := renderedLines[i-row][j-col]
				if char != ' ' {
					temp[j] = char
				}
			}
			tempStr := string(temp)
			if len(tempStr) < col+len(renderedLines[i-row]) {
				tempStr += renderedLines[i-row][col+len(renderedLines)-len(tempStr)-1:]
			}
			containerLines[i] = string(tempStr)
		}
	}

	containerRendered = string(strings.Join(containerLines, "\n"))

	if c.options.fixedSize.active {
		var tempContainer = make([]string, 0)
		for idx, line := range strings.Split(containerRendered, "\n") {
			if len(tempContainer) == int(c.options.fixedSize.height) {
				break
			}

			if len(line) < int(c.options.fixedSize.width) {
				tempContainer = append(tempContainer, line)
				for len(tempContainer[idx]) < int(c.options.fixedSize.width) {
					tempContainer[idx] += " "
				}
			} else {
				tempContainer = append(tempContainer, line[0:c.options.fixedSize.width])
			}
		}

		for len(tempContainer) != int(c.options.fixedSize.height) {
			tempContainer = append(tempContainer, strings.Repeat(" ", int(c.options.fixedSize.width)))
		}

		containerRendered = strings.Join(tempContainer, "\n")
	}

	return containerRendered
}
