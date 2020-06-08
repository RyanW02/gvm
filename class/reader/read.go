package reader

import (
	"encoding/binary"
	"errors"
)

const Magic = 0xCAFEBABE

var ErrUnderflow = errors.New("tried to read more bytes than present")
var ErrInvalidMagic = errors.New("magic didn't have value of 0xCAFEBABE")

func (cr *ClassReader) Read(n int) ([]byte, error) {
	bytes := make([]byte, n)

	read, err := cr.reader.Read(bytes)
	if err != nil {
		return nil, err
	}

	if read != n {
		return nil, ErrUnderflow
	}

	return bytes, nil
}

func (cr *ClassReader) ReadUint8() (uint8, error) {
	bytes, err := cr.Read(8 >> 3)
	if err != nil {
		return 0, nil
	}

	return bytes[0], nil
}

func (cr *ClassReader) ReadUint16() (uint16, error) {
	bytes, err := cr.Read(16 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint16(bytes), nil
}

func (cr *ClassReader) ReadUint32() (uint32, error) {
	bytes, err := cr.Read(32 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint32(bytes), nil
}

func (cr *ClassReader) ReadUint64() (uint64, error) {
	bytes, err := cr.Read(64 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint64(bytes), nil
}
