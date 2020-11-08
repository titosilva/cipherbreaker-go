package caesarbreaker

// CaesarBreaker struct
type CaesarBreaker struct {
	CipherText               string
	KeyList                  []byte
	BreakerThreadNumber      int
	TextAnalysisThreadNumber int
}
