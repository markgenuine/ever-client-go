package util

import "C"
import "encoding/hex"

type (
	// ExportedCChar ...
	ExportedCChar C.char

	// ExportedCInt ...
	ExportedCInt C.int
)

// CToString ...
func CToString(valueString *ExportedCChar, valueLen ExportedCInt) string {
	return C.GoStringN((*C.char)(valueString), (C.int)(valueLen))
}

// FromHex ...
func FromHex(value string) []byte {
	src := []byte(value)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, _ := hex.Decode(dst, src)

	return dst[:n]
}

//ToHex ...
func ToHex(value interface{}) []byte {
	switch date := value.(type) {
	case string:
		return byteToHex([]byte(date))
	case []byte:
		return byteToHex(date)
	}

	return []byte("")
}

func byteToHex(sl []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(sl)))
	hex.Encode(dst, sl)
	return dst
}
