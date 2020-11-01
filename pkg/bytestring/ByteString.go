package bytestring

// ByteString struct
// Define um tipo comum a ser consumido
// por todas as funções de criptografia
type ByteString struct {
	// String que poderá ser usada pelos encoders
	// Como forma de indicar algo sobre o formato
	Format string
	// Array de bytes para armazenar o valor
	Bytes []byte
}

// Invalid ByteString factory
func Invalid() ByteString {
	return ByteString{Format: "invalid", Bytes: nil}
}
