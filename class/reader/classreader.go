package reader

import "io"

// A ClassReader parses a single class
type ClassReader struct {
	reader io.Reader
}

func NewClassParser(reader io.Reader) ClassReader {
	return ClassReader{
		reader: reader,
	}
}
