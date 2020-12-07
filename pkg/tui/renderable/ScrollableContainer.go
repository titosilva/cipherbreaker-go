package renderable

import (
	"strings"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// ScrollableContainer ...
// Defines a container that can interact and be scrolled
type ScrollableContainer struct {
	Object
	InternalContainer Container
	scrolly           int
	width             int
	height            int
	killed            bool
}

// NewScrollableContainer function
// Constructor for ScrollableContainer
func NewScrollableContainer(width, height int) ScrollableContainer {
	return ScrollableContainer{
		InternalContainer: NewContainer(),
		scrolly:           0,
		width:             width,
		height:            height,
		killed:            false,
	}
}

// SetSize method
func (s *ScrollableContainer) SetSize(width, height int) {
	s.width = width
	s.height = height
}

// GetSize method
func (s *ScrollableContainer) GetSize() (width, height int) {
	return s.width, s.height
}

// SetScroll method
func (s *ScrollableContainer) SetScroll(y int) {
	if y < 0 {
		s.scrolly = 0
	} else {
		s.scrolly = y
	}
}

// GetScroll method
func (s *ScrollableContainer) GetScroll() (y int) {
	return s.scrolly
}

// ScrollDown method
func (s *ScrollableContainer) ScrollDown() {
	renderedHeight := len(strings.Split(s.InternalContainer.Render(), "\n"))
	if s.scrolly+s.height >= renderedHeight {
		if renderedHeight < s.height {
			s.scrolly = 0
		} else {
			s.scrolly = renderedHeight - s.height
		}
	} else {
		s.scrolly++
	}
}

// ScrollUp method
func (s *ScrollableContainer) ScrollUp() {
	if s.scrolly > 0 {
		s.scrolly--
	} else {
		s.scrolly = 0
	}
}

// ScrollUpMax method
func (s *ScrollableContainer) ScrollUpMax() {
	s.scrolly = 0
}

// ScrollDownMax method
func (s *ScrollableContainer) ScrollDownMax() {
	renderedHeight := len(strings.Split(s.InternalContainer.Render(), "\n"))
	if renderedHeight < s.height {
		s.scrolly = 0
	} else {
		s.scrolly = renderedHeight - s.height
	}
}

// Render method of ScrollableContainer
func (s *ScrollableContainer) Render() string {
	rendered := s.InternalContainer.Render()

	lines := make([]string, 0)
	min := s.scrolly + s.height
	renderedLines := strings.Split(rendered, "\n")

	if len(renderedLines) < min {
		min = len(renderedLines)
	}

	renderedLines = renderedLines[s.scrolly:min]
	newline := ""
	for _, line := range renderedLines {
		min = s.width

		if len(line) < min {
			min = len(line)
		}

		newline = line[0:min]
		lines = append(lines, newline)
	}

	for len(lines) < s.height {
		lines = append(lines, "")
	}

	newLines := make([]string, len(lines))
	for idx, line := range lines {
		newLines[idx] = line
		for len(newLines[idx]) < s.width {
			newLines[idx] += " "
		}
	}

	return strings.Join(newLines, "\n")
}

// Interact method of ScrollableContainer
func (s *ScrollableContainer) Interact() byte {
	for {
		select {
		case input := <-screen.InputChannel:
			switch input {
			case 'w':
				s.ScrollUp()
			case 's':
				s.ScrollDown()
			default:
				return input
			}
		}
	}
}

// DynamicRender method of ScrollableContainer
func (s *ScrollableContainer) DynamicRender(update chan bool) {
	go s.InternalContainer.DynamicRender(update)

	currentScrolly := s.scrolly
	for !s.killed {
		if s.scrolly != currentScrolly {
			update <- true
		}

		currentScrolly = s.scrolly
		time.Sleep(screen.RefreshMinDelay)
	}
}

// Kill method of ScrollableContainer
func (s *ScrollableContainer) Kill() {
	s.killed = true
}
