package attribute

import "github.com/RyanW02/gvm/class/reader"

type LocalVariableTypeTableEntry struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16
	SignatureIndex uint16
	Index          uint16
}

func NewLocalVariableTypeTableEntry() LocalVariableTypeTableEntry {
	return LocalVariableTypeTableEntry{}
}

func (e *LocalVariableTypeTableEntry) Read(cr *reader.ClassReader) (err error) {
	e.StartPc, err = cr.ReadUint16()
	e.Length, err = cr.ReadUint16()
	e.NameIndex, err = cr.ReadUint16()
	e.SignatureIndex, err = cr.ReadUint16()
	e.Index, err = cr.ReadUint16()
	return
}
