package multiwordmatcher

import "testing"

type MultiWordMatcherTestCase struct {
	text     string
	expected []MatchInfo
}

func genTestCases() []MultiWordMatcherTestCase {
	cases := make([]MultiWordMatcherTestCase, 0)

	cases = append(cases, MultiWordMatcherTestCase{
		text: "o AbacAxi e o MorAngo sao frutas diferentes",
		expected: []MatchInfo{
			{
				Word:    "abacaxi",
				Indexes: []int{2},
			},
			{
				Word:    "morango",
				Indexes: []int{14},
			},
		},
	})

	cases = append(cases, MultiWordMatcherTestCase{
		text: "o abacaxi e o morango sao frutas diferentes",
		expected: []MatchInfo{
			{
				Word:    "abacaxi",
				Indexes: []int{2},
			},
			{
				Word:    "morango",
				Indexes: []int{14},
			},
		},
	})

	cases = append(cases, MultiWordMatcherTestCase{
		text: "como o abacaxi, degusto o abacaxi",
		expected: []MatchInfo{
			{
				Word:    "abacaxi",
				Indexes: []int{7, 26},
			},
		},
	})

	cases = append(cases, MultiWordMatcherTestCase{
		text: "como o abacaxi, degusto o mOrango",
		expected: []MatchInfo{
			{
				Word:    "abacaxi",
				Indexes: []int{7},
			},
			{
				Word:    "morango",
				Indexes: []int{26},
			},
		},
	})

	return cases
}

func TestMultiWordMatcher(t *testing.T) {
	dry := Dry{
		WordList:      []string{"abacaxi", "morango"},
		ThreadsNumber: 1,
	}

	for _, testCase := range genTestCases() {
		matchInfo := dry.Analyse(testCase.text)
		testMatch(matchInfo, testCase.expected, t)
	}
}

func testMatch(matchInfo []MatchInfo, expected []MatchInfo, t *testing.T) {
	if len(matchInfo) != len(expected) {
		t.Errorf("MultiWordMatcher FAILED -> expected %d matches, but found %d", len(expected), len(matchInfo))
	}

	for _, match := range matchInfo {
		ok := false
		for _, expectedMatch := range expected {
			if match.Word == expectedMatch.Word {
				ok = true
				// Compare the list of indexes
				if len(match.Indexes) != len(expectedMatch.Indexes) {
					t.Errorf("MultiWordMatcher FAILED -> for word %s, expected %d indexes, but found %d", match.Word, len(expectedMatch.Indexes),
						len(match.Indexes))
					ok = false
				} else {
					for _, matchIndex := range match.Indexes {
						found := false
						for _, expectedIndex := range expectedMatch.Indexes {
							if matchIndex == expectedIndex {
								found = true
								break
							}
						}

						if !found {
							t.Errorf("MultiWordMatcher FAILED -> for word %s, found index %d that was not expected", match.Word, matchIndex)
							ok = false
							break
						}
					}
				}
			}
		}

		if !ok {
			t.Errorf("MultiWordMatcher FAILED -> expected: %v; got: %v", expected, matchInfo)
			break
		}
	}
}
