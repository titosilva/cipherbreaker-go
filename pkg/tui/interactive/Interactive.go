package interactive

import "github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"

// UserInteractive interface
// Defines a interface for a entity
// that may interact with the user in some way
// using the Interact method (it will "take control" of the interaction)
type UserInteractive interface {
	renderable.DynamicRenderable
	Interact(args ...interface{}) byte
}
