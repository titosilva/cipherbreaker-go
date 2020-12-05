package section

// Section interface
// An interface is a part of the running
// program. It takes the control of the executio
// through the method Run, and returns the next Section
// to be run
type Section interface {
	Run(next *Section)
}
