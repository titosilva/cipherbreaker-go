package renderable

import (
	"fmt"
	"strings"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// OptionsList struct
type OptionsList struct {
	Object
	optionList    []Renderable
	selected      int
	selectionChar byte
	killed        bool
}

// NewOptionsList funcion
// Constructor for OptionsList
func NewOptionsList(selectionChar byte) OptionsList {
	return OptionsList{optionList: make([]Renderable, 0), selected: 0, selectionChar: selectionChar}
}

// AddOption method of OptionList
// Adds a renderable object to the list
func (oplist *OptionsList) AddOption(option Renderable) {
	oplist.optionList = append(oplist.optionList, option)
}

// GetSelectedOption method of OptionsList
// Returns the selected option as an int
func (oplist *OptionsList) GetSelectedOption() int {
	return oplist.selected
}

// Render method of OptionsList
func (oplist *OptionsList) Render() string {
	c := NewContainer()

	lineCounter := 0
	selectionLine := 0
	for idx, item := range oplist.optionList {
		item.SetPosition(uint(lineCounter), 3)
		c.AddItem(item)
		itemHeigth := len(strings.Split(item.Render(), "\n"))
		// determines the selection position
		if idx == oplist.selected {
			selectionLine = lineCounter + (itemHeigth / 2)
		}
		lineCounter += itemHeigth
	}

	// Add the selection char in the selection line
	text := NewText(fmt.Sprintf(" %c ", oplist.selectionChar), uint(selectionLine), 0)
	c.AddItem(&text)

	return c.Render()
}

// DynamicRender method of OptionsList
// Keeps checking for changes in the selected item
func (oplist *OptionsList) DynamicRender(update chan bool) {
	currentSelection := oplist.selected
	for !oplist.killed {
		// Check if selection has changed
		if oplist.selected != currentSelection {
			// If yes, request update
			update <- true
		}

		currentSelection = oplist.selected
		time.Sleep(screen.RefreshMinDelay)
	}
}

// Kill method of OptionsList
func (oplist *OptionsList) Kill() {
	oplist.killed = true
}

// Interact method of OptionsList
// Interacts with the user to select an option
func (oplist *OptionsList) Interact() byte {
	numOfOptions := len(oplist.optionList)
	for {
		select {
		case input := <-screen.InputChannel:
			switch input {
			case 's':
				oplist.selected = (oplist.selected + 1) % numOfOptions
			case 'w':
				oplist.selected = (oplist.selected - 1) % numOfOptions
			case screen.KeyEscape, screen.KeyEnter:
				return input
			default:
				time.Sleep(screen.RefreshMinDelay)
			}
		default:
			time.Sleep(screen.RefreshMinDelay)
		}
	}
}
