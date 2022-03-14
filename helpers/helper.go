package helpers

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(number int) ([]byte, error) {
	buff := bytes.Buffer{}

	err := binary.Write(&buff, binary.BigEndian, int64(number))

	return buff.Bytes(), err
}
