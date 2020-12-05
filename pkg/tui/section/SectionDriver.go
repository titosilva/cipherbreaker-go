package section

// StartExecution function
// Runs the Sections automatically going to
// the next Section till it receives a nil as
// next section
func StartExecution(start Section) {
	var next Section
	var current Section

	start.Run(&next)

	for next != nil {
		current = next
		current.Run(&next)
	}
}
