package vigenere

import "fmt"

// Vigenere cipher base type
type Vigenere struct{}

// Vigenere's cipher implementation.

// Cipher method for Vigenere's Cipher
// Takes arguments as string slices to enforce use of ASCII encoding
// (golang doesn't have a char[] type, but has byte[] and rune[])
func (c Vigenere) Cipher(plainTextString string, key string) (string, error) {
	plainText := []byte(plainTextString)
	cipherText := make([]byte, len(plainText))
	keyAlpha := make([]byte, len(key))

	// Standardizing the format of the key
	for i := range key {
		if (key[i] >= 'a') && (key[i]) <= 'z' {
			keyAlpha[i] = key[i] - 'a'
		} else if (key[i] >= 'A') && (key[i] <= 'Z') {
			keyAlpha[i] = key[i] - 'A'
		} else {
			return "", fmt.Errorf("Key %s is not valid. Use [a-z] or [A-Z]", key)
		}
	}

	//Encoding the plainText into the cipherText using the keyAlpha
	idxKey := 0
	for idx := range plainText {
		if plainText[idx] >= 'a' && plainText[idx] <= 'z' {
			cipherText[idx] = ((plainText[idx]-'a')+keyAlpha[idxKey])%26 + 'a'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else if plainText[idx] >= 'A' && plainText[idx] <= 'Z' {
			cipherText[idx] = ((plainText[idx]-'A')+keyAlpha[idxKey])%26 + 'A'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else {
			cipherText[idx] = plainText[idx]
		}
	}

	return string(cipherText), nil
}

// Vigenere's Decipher implementation

// Decipher method for Vigenere's Cipher
// Takes arguments as string slices to enforce use of ASCII encoding
// (golang doesn't have a char[] type, but has byte[] and rune[])
func (c Vigenere) Decipher(cipherTextString string, key string) (string, error) {
	cipherText := []byte(cipherTextString)
	decipherText := make([]byte, len(cipherText))
	keyAlpha := make([]byte, len(key))

	// Standardizing the format of the key
	for i := range key {
		if (key[i] >= 'a') && (key[i]) <= 'z' {
			keyAlpha[i] = key[i] - 'a'
		} else if (key[i] >= 'A') && (key[i] <= 'Z') {
			keyAlpha[i] = key[i] - 'A'
		} else {
			return "", fmt.Errorf("Key %s is not valid. Use [a-z] or [A-Z]", key)
		}
	}

	//Encoding the cipherText into the decipherText using the keyAlpha
	idxKey := 0
	for idx := range cipherText {
		if cipherText[idx] >= 'a' && cipherText[idx] <= 'z' {
			decipherText[idx] = ((cipherText[idx]-'a')+26-keyAlpha[idxKey])%26 + 'a'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else if cipherText[idx] >= 'A' && cipherText[idx] <= 'Z' {
			decipherText[idx] = ((cipherText[idx]-'A')+26-keyAlpha[idxKey])%26 + 'A'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else {
			decipherText[idx] = cipherText[idx]
		}
	}

	return string(decipherText), nil
}
