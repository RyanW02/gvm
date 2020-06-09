package attribute

import "github.com/RyanW02/gvm/class/reader"

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}

func NewLineNumberTableEntry() LineNumberTableEntry {
	return LineNumberTableEntry{}
}

func (e *LineNumberTableEntry) Read(cr *reader.ClassReader) (err error) {
	e.StartPc, err = cr.ReadUint16()
	e.LineNumber, err = cr.ReadUint16()
	return
}
