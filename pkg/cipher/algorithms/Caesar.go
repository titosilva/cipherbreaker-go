package algorithms

import "fmt"

// Caesar cipher base type
type Caesar struct{}

// Cipher method for Caesar Cipher
// Takes arguments as byte slices to enforce use of ASCII encoding
// (golang doesn't have a char type, but has byte and rune)
func (c Caesar) Cipher(plainTextString string, key byte) (string, error) {
	plainText := []byte(plainTextString)
	cipherText := make([]byte, len(plainText))
	var keyAlpha byte

	switch {
	case 'a' <= key && key <= 'z':
		keyAlpha = key - 'a'
	case 'A' <= key && key <= 'Z':
		keyAlpha = key - 'A'
	default:
		return "", fmt.Errorf("Key %c is not valid. Use [a-z] or [A-Z]", key)
	}

	for idx, char := range plainText {
		if 'a' <= char && char <= 'z' {
			cipherText[idx] = char + (keyAlpha)
		} else if 'A' <= char && char <= 'Z' {
			cipherText[idx] = char + (keyAlpha)
		} else {
			cipherText[idx] = char
		}
	}

	return string(cipherText), nil
}
