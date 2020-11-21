package caesar

import (
	"testing"
)

// Test Caesar cipher method
func TestCaesarCipher(t *testing.T) {
	cases := []struct {
		caseString string
		caseKey    byte
		expected   string
		// Tells if a case is of success or fail
		success bool
	}{
		{
			caseString: "abc",
			caseKey:    'b',
			expected:   "bcd",
			success:    true,
		},
		{
			caseString: "abc 123",
			caseKey:    'b',
			expected:   "bcd 123",
			success:    true,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    'e',
			expected:   "Himr Lyrh mwx wilv wxevo.",
			success:    true,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    'Z',
			expected:   "Cdhm Gtmc hrs rdgq rszqj.",
			success:    true,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    '9',
			success:    false,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    '~',
			success:    false,
		},
	}

	for _, c := range cases {
		if c.success {
			// Success cases
			t.Logf("Caesar testing: %s <key: %c> -> %s", c.caseString, c.caseKey, c.expected)
			result, err := Cipher(c.caseString, c.caseKey)

			if err != nil {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected: %s; got ERROR: %s", c.caseString, c.caseKey, c.expected, err)
			}

			if result != c.expected {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected: %s; got: %s", c.caseString, c.caseKey, c.expected, result)
			}
		} else {
			// Fail cases
			t.Logf("Caesar testing: %s <key: %c> -> expected err", c.caseString, c.caseKey)
			result, err := Cipher(c.caseString, c.caseKey)

			if err == nil {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected error, but got: %s", c.caseString, c.caseKey, result)
			}
		}
	}
}

// Test Caesar cipher method
func TestCaesarDipher(t *testing.T) {
	cases := []struct {
		caseString string
		caseKey    byte
		expected   string
		// Tells if a case is of success or fail
		success bool
	}{
		{
			caseString: "bcd",
			caseKey:    'b',
			expected:   "abc",
			success:    true,
		},
		{
			expected:   "abc 123",
			caseKey:    'b',
			caseString: "bcd 123",
			success:    true,
		},
		{
			expected:   "Dein Hund ist sehr stark.",
			caseKey:    'e',
			caseString: "Himr Lyrh mwx wilv wxevo.",
			success:    true,
		},
		{
			caseString: "Cdhm Gtmc hrs rdgq rszqj.",
			caseKey:    'Z',
			expected:   "Dein Hund ist sehr stark.",
			success:    true,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    '9',
			success:    false,
		},
		{
			caseString: "Dein Hund ist sehr stark.",
			caseKey:    '~',
			success:    false,
		},
	}

	for _, c := range cases {
		if c.success {
			// Success cases
			t.Logf("Caesar testing: %s <key: %c> -> %s", c.caseString, c.caseKey, c.expected)
			result, err := Decipher(c.caseString, c.caseKey)

			if err != nil {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected: %s; got ERROR: %s", c.caseString, c.caseKey, c.expected, err)
			}

			if result != c.expected {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected: %s; got: %s", c.caseString, c.caseKey, c.expected, result)
			}
		} else {
			// Fail cases
			t.Logf("Caesar testing: %s <key: %c> -> expected err", c.caseString, c.caseKey)
			result, err := Decipher(c.caseString, c.caseKey)

			if err == nil {
				t.Errorf("Caesar FAILED: %s <key: %c> -> expected error, but got: %s", c.caseString, c.caseKey, result)
			}
		}
	}
}
