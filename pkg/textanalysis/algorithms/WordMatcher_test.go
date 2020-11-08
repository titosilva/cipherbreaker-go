package algorithms

import "testing"

func TestWordMatcher(t *testing.T) {
	testcases := []struct {
		// Word to match
		caseWord string
		// String to test
		caseStr string
		// Expected response
		expected bool
		// Indicates if the test case is a
		// success or a fail case
		success bool
	}{
		{
			caseWord: "abacaxi",
			caseStr:  "O abadaxi estava muito bom. Ops, quero dizer, abacaxi",
			expected: true,
			success:  true,
		}, {
			caseWord: "abacaxi",
			caseStr:  "O abadaxi estava muito bom.",
			expected: false,
			success:  true,
		}, {
			// Test case sensitibility - must be case insensitive
			caseWord: "Hund",
			caseStr:  "Dein hund ist sehr klug",
			expected: true,
			success:  true,
		},
	}

	// iterate over test cases
	for _, c := range testcases {
		if c.success {
			// Success cases
			match, err := WordMatcher{Word: c.caseWord}.Analyse(c.caseStr)

			if err != nil {
				t.Errorf("WordMatcher FAILED: %s -> expected: %t; got ERROR: %s", c.caseWord, c.expected, err)
			}

			if match != c.expected {
				t.Errorf("WordMatcher FAILED: %s -> expected: %t; got: %t", c.caseWord, c.expected, match)
			}
		} else {
			// Fail cases
			t.Log("Ok")
		}
	}
}
