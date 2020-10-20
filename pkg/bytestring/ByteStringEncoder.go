package bytestring

// Encoder interface
// Interface para os encoders
// encoders são estruturas com capacidade transformar
// algum tipo de dado em um ByteString para uso posterior
// Vários encoders podem partir de um mesmo tipo de dado,
// e substituir apenas o algoritmo utilizado
type Encoder interface {
	// Método de codificação
	Encode() ByteString
}
