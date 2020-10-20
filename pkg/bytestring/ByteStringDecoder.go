package bytestring

// Decoder interface
// Interface para os decoders
// decoders são estruturas com capacidade transformar
// ByteString em outros tipos de dado para uso posterior
// Decoders diferentes podem gerar um mesmo tipo, trocando
// apenas o algoritmo utilizado
type Decoder interface {
	// Método de codificação
	Decode(ByteString)
}
