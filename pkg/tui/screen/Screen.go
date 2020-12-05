package screen

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/term"
)

const (
	// RefreshMinDelay -> time to wait on request updates
	RefreshMinDelay = 100 * time.Millisecond
)

var (
	inputActive  bool
	InputChannel chan byte
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

	go func() {
		inputActive = true
		InputChannel = make(chan byte, 100)
		for inputActive {

			input, _ := ReadByte()

			InputChannel <- input
		}
		close(InputChannel)
	}()

	return state, nil
}

// Restore function
// Set the terminal to old state
func Restore(state *term.State) {
	inputActive = false
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
	lines := strings.Split(content, "\n")
	for _, line := range lines[:len(lines)-1] {
		fmt.Printf(line + "\r\n")
	}
	fmt.Printf(lines[len(lines)-1])
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
