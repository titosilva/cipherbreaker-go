package main

import (
	"github.com/titosilva/cipherbreaker-go/internal/cipherbreakersections"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/screen"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
)

func main() {
	oldState, err := screen.Init()

	if err != nil {
		println("Failed to configure terminal :(")
		println(err)
		return
	}

	defer screen.Restore(oldState)
	var main cipherbreakersections.MainSection
	section.StartExecution(main)
}
