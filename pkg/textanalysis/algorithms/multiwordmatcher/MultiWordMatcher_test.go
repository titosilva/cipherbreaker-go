package multiwordmatcher

import "testing"

func TestMultiWordMatcher(t *testing.T) {
	dry := Dry{
		WordList:      []string{"abacaxi", "morango"},
		ThreadsNumber: 3,
	}

	matchInfo := dry.Analyse("o Abacaxi e o Morango sao frutas diferentes")

	expected := []MatchInfo{
		{
			Word:    "abacaxi",
			Indexes: []int{2},
		},
		{
			Word:    "morango",
			Indexes: []int{14},
		},
	}

	for _, match := range matchInfo {
		ok := false
		for _, expectedMatch := range expected {
			if match.Word == expectedMatch.Word {
				ok = true
				for _, matchIndexes := range match.Indexes {
					for _, expectedIndexes := range expectedMatch.Indexes {
						if matchIndexes == expectedIndexes {
							break
						}
					}
				}
			}
		}

		if !ok {
			t.Errorf("MultiWordMatcher failed-> expected: %v; got: %v", matchInfo, expected)
			break
		}
	}
}
