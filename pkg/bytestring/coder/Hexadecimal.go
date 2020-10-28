package coder

import (
	"fmt"

	"github.com/titosilva/cipherbreaker-go/pkg/bytestring"
)

// HexStringEncoder struct
type HexStringEncoder struct {
	HexString string
}

// Encode ...
func (hexencoder HexStringEncoder) Encode() bytestring.ByteString {
	fmt.Println("Ok")
	return bytestring.ByteString{Format: "hex", Bytes: []byte{0}}
}
