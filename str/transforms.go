package str

import (
	"bytes"
)

func Title(b []byte) ([]byte, error) {
	b = bytes.ToLower(b)
	return bytes.Title(b), nil
}
