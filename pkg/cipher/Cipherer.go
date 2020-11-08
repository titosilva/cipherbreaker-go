package cipher

// Cipherer interface
// Takes any arguments (to be decided by the implementation) and
// makes an encryption
type Cipherer interface {
	Cipher(args ...interface{}) interface{}
}
