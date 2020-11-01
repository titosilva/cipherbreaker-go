package coder

import (
	"fmt"
	"testing"

	"github.com/titosilva/cipherbreaker-go/pkg/bytestring"
)

// TestCase struct
type TestCase struct {
	caseStr    string
	expectedBS bytestring.ByteString
}

// GenTestCases Function
// Generates test cases for a given format and the corresponding
// expected bytestrings
func GenTestCases() []TestCase {
	var testcases = make([]TestCase, 0)
	for i := 0; i < 10; i++ {
		b := byte(i) << 4
		testcases = append(testcases, TestCase{
			caseStr:    fmt.Sprintf("%d", i),
			expectedBS: bytestring.ByteString{Format: "hex", Bytes: []byte{b}},
		})
	}

	for i := 0; i < 6; i++ {
		c := 'a' + i
		b := byte(10+i) << 4

		testcases = append(testcases, TestCase{
			caseStr:    fmt.Sprintf("%c", c),
			expectedBS: bytestring.ByteString{Format: "hex", Bytes: []byte{b}},
		})
	}

	testcases = append(testcases, TestCase{
		caseStr: "aabbcc1234",
		expectedBS: bytestring.ByteString{
			Format: "hex",
			Bytes: []byte{
				170,
				187,
				204,
				18,
				52,
			},
		},
	})

	return testcases
}

// TestHexStringEncoderEncode function
// Tests for HexStringEncoder encode function
func TestHexStringEncoderEncode(t *testing.T) {
	for _, hexcase := range GenTestCases() {
		bs, err := HexStringEncoder{String: hexcase.caseStr}.Encode()

		if err != nil {
			t.Errorf("HexStringEncoder FAILED -> conversion gave error: %s", err)
		}

		if bs.Format != "hex" {
			t.Errorf("HexStringEncoder FAILED -> case: %s; format: %s; expected: %s", hexcase.caseStr, bs.Format, hexcase.expectedBS.Format)
		} else {
			for i, b := range bs.Bytes {
				if b != hexcase.expectedBS.Bytes[i] {
					t.Errorf("HexStringEncoder FAILED -> case: %s; got: %+v; expect: %+v", hexcase.caseStr, bs, hexcase.expectedBS)
					break
				}
			}
		}
	}
}
