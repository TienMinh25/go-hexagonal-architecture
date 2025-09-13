package encoding

import (
	"encoding/hex"

	"aidanwoods.dev/go-result/result"
)

// Encode hex
func HexEncode(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// Decode hex
func HexDecode(encoded string) result.Result[[]byte] {
	if b, err := hex.DecodeString(encoded); err != nil {
		return result.Err[[]byte](err)
	} else {
		return result.Ok(b)
	}
}
