package renderable

// DynamicRenderable interface
// Defines an interface to an entity
// that may change at runtime, and request
// an screen update to show its new contents
type DynamicRenderable interface {
	DynamicRender(update chan bool)
	Kill()
}
