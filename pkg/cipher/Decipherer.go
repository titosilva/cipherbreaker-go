package cipher

// Decipherer interface
// Takes any arguments (to be decided by the implementation) and
// makes an decryption
type Decipherer interface {
	Decipher(args ...interface{}) interface{}
}
