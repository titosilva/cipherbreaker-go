package section

// StartExecution function
// Runs the Sections automatically going to
// the next Section till it receives a nil as
// next section
func StartExecution(start Section) {
	var next Section

	next = start.Run()
	for !next.IsEnd() {
		next = next.Run()
	}
}
