package utils

import "encoding/hex"

func EncodeHex(b []byte) string {
	return hex.EncodeToString(b)
}

func DecodeHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func MustDecodeHex(s string) []byte {
	bt, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bt
}
