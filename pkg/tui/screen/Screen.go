package screen

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.org/x/term"
)

const (
	// RefreshMinDelay -> time to wait on request updates
	RefreshMinDelay = 100 * time.Millisecond
)

// Init function of screen
// Sets the terminal to raw mode
// and returns the state to be restored
// later
// Returns error if it can't get the state of Stdin
func Init() (*term.State, error) {
	state, err := term.GetState(int(os.Stdin.Fd()))

	if err != nil {
		return nil, err
	}

	term.MakeRaw(int(os.Stdin.Fd()))
	return state, nil
}

// Restore function
// Set the terminal to old state
func Restore(state *term.State) {
	term.Restore(int(os.Stdin.Fd()), state)
}

// GetSize function
// Returns the size of the screen (terminal)
func GetSize() (width, height int) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		width = 200
		height = 40
	}

	return width, height
}

// Clear function
// clears the screen
func Clear() {
	cmd := exec.Command("clear || cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Print function
// Prints on the screen
func Print(content string) {
	fmt.Print(content)
}

// ReadByte function
// Reads a byte from input
func ReadByte() (byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadByte()

	if err != nil {
		return 0, err
	}

	return input, nil
}
