package renderable

import (
	"fmt"
	"strings"
	"time"

	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
)

// OptionsList struct
type OptionsList struct {
	optionList    []Renderable
	selected      int
	selectionChar byte
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
func (oplist *OptionsList) DynamicRender(update chan bool) {
	currentSelection := oplist.selected
	for true {
		// Check if selection has changed
		if oplist.selected != currentSelection {
			// If yes, request update
			update <- true
		}

		currentSelection = oplist.selected
		time.Sleep(screen.RefreshMinDelay)
	}
}

// Interact method of OptionsList
// Interacts with the user to select an option
func (oplist *OptionsList) Interact() {
	input, err := screen.ReadByte()

	if err != nil {
		return
	}

	switch input {
	case 'd':

	default:
	}
}
