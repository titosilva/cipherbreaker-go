package cipherbreaker

// CipherBreaker interface
// Simple interface to be used to implement cipher breakers
type CipherBreaker interface {
	Break(args ...interface{}) interface{}
}
