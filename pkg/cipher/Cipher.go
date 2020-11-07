package cipher

// Cipher interface
// Implements both Cipherer and Decipherer interfaces
type Cipher interface {
	Cipherer
	Decipherer
}
