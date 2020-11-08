package wordmatcher

import "testing"

func TestWordMatcher(t *testing.T) {
	testcases := []struct {
		// Word to match
		caseWord string
		// String to test
		caseStr string
		// Expected response
		expectedMatch bool
		expectedIdx   int
		// Indicates if the test case is a
		// success or a fail case
		success bool
	}{
		{
			caseWord:      "abacaxi",
			caseStr:       "O abadaxi estava muito bom. Ops, quero dizer, abacaxi",
			expectedMatch: true,
			expectedIdx:   46,
			success:       true,
		}, {
			caseWord:      "abacaxi",
			caseStr:       "O abadaxi estava muito bom.",
			expectedMatch: false,
			expectedIdx:   -1,
			success:       true,
		}, {
			// Test case sensitibility - must be case insensitive
			caseWord:      "Hund",
			caseStr:       "Dein hund ist sehr klug",
			expectedMatch: true,
			expectedIdx:   5,
			success:       true,
		},
	}

	// iterate over test cases
	for _, c := range testcases {
		if c.success {
			// Success cases
			match, idx := WordMatcher{Word: c.caseWord}.Analyse(c.caseStr)

			if idx != c.expectedIdx {
				t.Errorf("WordMatcher FAILED: %s -> expected index: %d; got index: %d", c.caseWord, c.expectedIdx, idx)
			}

			if match != c.expectedMatch {
				t.Errorf("WordMatcher FAILED: %s -> expected: %t; got: %t", c.caseWord, c.expectedMatch, match)
			}
		} else {
			// Fail cases
			t.Log("Ok")
		}
	}
}
