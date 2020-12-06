package cipherbreakersections

import "github.com/titosilva/cipherbreaker-go/pkg/tui/section"

// EndSection struct
// Ends the execution
type EndSection struct{}

// IsEnd method of EndSection
// Only returns true
func (end EndSection) IsEnd() bool {
	return true
}

// Run method of EndSection
// Does nothing
func (end EndSection) Run() section.Section { return EndSection{} }
