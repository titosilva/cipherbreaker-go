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
	var keyAlpha []byte

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

	//Equalling the size of the key with the text, if the key is smaller than the text
	for i, aux := 0, 0; i < len(plainText); i++ {
		if len(keyAlpha) < len(plainText) && len(keyAlpha) > 1 && keyAlpha[i] == 0 && plainText[i] != 0 {
			keyAlpha[i] = keyAlpha[aux]
			aux++
		} else if len(key) < len(plainText) {
			keyAlpha[i] = plainText[i]
		} else {
			break
		}
	}

	//Encoding the plainText into the cipherText using the keyAlpha
	for idx := range plainText {
		if cipherText[idx] >= 'a' && cipherText[idx] <= 'z' {
			cipherText[idx] = (plainText[idx]+keyAlpha[idx])%26 + 'a'
		} else if cipherText[idx] >= 'A' && cipherText[idx] <= 'Z' {
			cipherText[idx] = (plainText[idx]+keyAlpha[idx])%26 + 'A'
		} else {
			cipherText[idx] = keyAlpha[idx]
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
	var keyAlpha []byte

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

	//Equalling the size of the key with the text, if the key is smaller than the text
	for i, aux := 0, 0; i < len(cipherText); i++ {
		if len(keyAlpha) < len(cipherText) && len(keyAlpha) > 1 && keyAlpha[i] == 0 && cipherText[i] != 0 {
			keyAlpha[i] = keyAlpha[aux]
			aux++
		} else if len(key) < len(cipherText) {
			keyAlpha[i] = cipherText[i]
		} else {
			break
		}
	}

	//Decoding the plainText into the cipherText using the keyAlpha
	for idx := range cipherText {
		if decipherText[idx] >= 'a' && decipherText[idx] <= 'z' {
			decipherText[idx] = (cipherText[idx]-keyAlpha[idx]+26)%26 + 'a'
		} else if decipherText[idx] >= 'A' && decipherText[idx] <= 'Z' {
			decipherText[idx] = (cipherText[idx]-keyAlpha[idx]+26)%26 + 'A'
		} else {
			decipherText[idx] = keyAlpha[idx]
		}
	}
	return string(decipherText), nil
}
