package attribute

import "github.com/RyanW02/gvm/class/reader"

type LocalVariableTableEntry struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

func NewLocalVariableTableEntry() LocalVariableTableEntry {
	return LocalVariableTableEntry{}
}

func (e *LocalVariableTableEntry) Read(cr *reader.ClassReader) (err error) {
	e.StartPc, err = cr.ReadUint16()
	e.Length, err = cr.ReadUint16()
	e.NameIndex, err = cr.ReadUint16()
	e.DescriptorIndex, err = cr.ReadUint16()
	e.Index, err = cr.ReadUint16()
	return
}
