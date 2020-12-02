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

		{
			caseString: "Fim de semestre eh assim",
			caseKey:    "hard",
			expected:   "Mid gl svplskul ey dzszp",
			success:    true,
		},

		{
			caseString: "this year was a tragic year",
			caseKey:    "corona",
			expected:   "vvzg lecf nof a vfruvc asrf",
			success:    true,
		},

		{
			caseString: "die Kunst des Rechnens",
			caseKey:    "GOLANG",
			expected:   "jwp Khtyh oef Xkqsnrty",
			success:    true,
		},

		{
			caseString: "a chave de codificacao eh restrita",
			caseKey:    "%",
			expected:   "? a,?:) () a3(-*-a?a?3 ), 6)786-8?",
			success:    false,
		},

		{
			caseString: "somente caracteres alfabeticos da ascii podem ser utilizados",
			caseKey:    "123",
			expected:   "c@?5?f5 43b25d6d5d 3<7326f94ac 53 1d59: b?57= d7b ff9=;j26?d",
			success:    false,
		},

		{
			caseString: "Porem, tanto faz Usar MaIUsCUlo ou MINUSculo",
			caseKey:    "GOisNice",
			expected:   "Vczwz, bcrzc nsm Cuex AiAHaEYrc wm ZQPYYqcdb",
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

		{
			caseString: "Mid gl svplskul ey dzszp",
			caseKey:    "hard",
			expected:   "Fim de semestre eh assim",
			success:    true,
		},

		{
			caseString: "vvzg lecf nof a vfruvc asrf",
			caseKey:    "corona",
			expected:   "this year was a tragic year",
			success:    true,
		},

		{
			caseString: "jwp Khtyh oef Xkqsnrty",
			caseKey:    "GOLANG",
			expected:   "die Kunst des Rechnens",
			success:    true,
		},

		{
			caseString: "? a,?:) () a3(-*-a?a?3 ), 6)786-8?",
			caseKey:    "%",
			expected:   "a chave de codificacao eh restrita",
			success:    false,
		},

		{
			caseString: "c@?5?f5 43b25d6d5d 3<7326f94ac 53 1d59: b?57= d7b ff9=;j26?d",
			caseKey:    "123",
			expected:   "somente caracteres alfabeticos da ascii podem ser utilizados",
			success:    false,
		},

		{
			caseString: "Vczwz, bcrzc nsm Cuex AiAHaEYrc wm ZQPYYqcdb",
			caseKey:    "GOisNice",
			expected:   "Porem, tanto faz Usar MaIUsCUlo ou MINUSculo",
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
