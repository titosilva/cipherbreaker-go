package textanalysis

// Analyzer interface
// Receives a text, and returns some information about that text
type Analyzer interface {
	Analyze(args ...interface{}) interface{}
}
