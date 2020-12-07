package caesarbreaker

// GenerateKeys function of CaesarBreaker module
// Sends all alphabetic bytes to the channel
func GenerateKeys(key chan byte) {
	for i := 0; i < 26; i++ {
		key <- byte('a' + i)
	}

	close(key)
}
