package vigenere

import (
	"testing"
)

// Test Vigenere cipher method
func TestVigenereCipher(t *testing.T) {
	var vigenere Vigenere

	cases := []struct {
		caseString string
		caseKey    string
		expected   string
		// Tells if a case is of success or fail
		success bool
	}{
		{
			caseString: "Deus e bom, o tempo todo",
			caseKey:    "UnB",
			expected:   "Xrvm r ciz, p nrnjb uiqp",
			success:    true,
		},
	}
	for _, c := range cases {
		if c.success {
			// Success cases
			t.Logf("Vigenere testing: %s <key: %s> -> %s", c.caseString, c.caseKey, c.expected)
			result, err := vigenere.Cipher(c.caseString, c.caseKey)

			if err != nil {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected: %s; got ERROR: %s", c.caseString, c.caseKey, c.expected, err)
			}

			if result != c.expected {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected: %s; got: %s", c.caseString, c.caseKey, c.expected, result)
			}
		} else {
			// Fail cases
			t.Logf("Vigenere testing: %s <key: %s> -> expected err", c.caseString, c.caseKey)
			result, err := vigenere.Cipher(c.caseString, c.caseKey)

			if err == nil {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected error, but got: %s", c.caseString, c.caseKey, result)
			}
		}
	}
}

func TestVigenereDecipher(t *testing.T) {
	var vigenere Vigenere

	cases := []struct {
		caseString string
		caseKey    string
		expected   string
		// Tells if a case is of success or fail
		success bool
	}{
		{
			caseString: "Xrvm r ciz, p nrnjb uiqp",
			caseKey:    "UnB",
			expected:   "Deus e bom, o tempo todo",
			success:    true,
		},
	}
	for _, c := range cases {
		if c.success {
			// Success cases
			t.Logf("Vigenere testing: %s <key: %s> -> %s", c.caseString, c.caseKey, c.expected)
			result, err := vigenere.Decipher(c.caseString, c.caseKey)

			if err != nil {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected: %s; got ERROR: %s", c.caseString, c.caseKey, c.expected, err)
			}

			if result != c.expected {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected: %s; got: %s", c.caseString, c.caseKey, c.expected, result)
			}
		} else {
			// Fail cases
			t.Logf("Vigenere testing: %s <key: %s> -> expected err", c.caseString, c.caseKey)
			result, err := vigenere.Decipher(c.caseString, c.caseKey)

			if err == nil {
				t.Errorf("Vigenere FAILED: %s <key: %s> -> expected error, but got: %s", c.caseString, c.caseKey, result)
			}
		}
	}
}
