package coder

import (
	"fmt"
	"strconv"

	"github.com/titosilva/cipherbreaker-go/pkg/bytestring"
)

// HexStringEncoder struct
type HexStringEncoder struct {
	String string
}

// Encode ...
func (hexencoder HexStringEncoder) Encode() (bs bytestring.ByteString, err error) {
	var newbs = make([]byte, 0)
	var b uint64
	var newerr error

	for i := 0; i < len(hexencoder.String); i += 2 {
		if i == len(hexencoder.String)-1 {
			s := fmt.Sprintf("%c%c", hexencoder.String[i], '0')
			b, newerr = strconv.ParseUint(s, 16, 8)
		} else {
			s := hexencoder.String[i : i+2]
			b, newerr = strconv.ParseUint(s, 16, 8)
		}

		if err != nil {
			return bytestring.Invalid(), newerr
		}

		newbs = append(newbs, byte(b))
	}

	return bytestring.ByteString{Format: "hex", Bytes: newbs}, nil
}
