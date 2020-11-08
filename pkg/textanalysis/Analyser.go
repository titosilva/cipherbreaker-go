package textanalysis

// Analyser interface
// Receives a text, and returns some information about that text
type Analyser interface {
	Analyse(string) interface{}
}
