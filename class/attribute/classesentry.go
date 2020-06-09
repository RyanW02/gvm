package attribute

import (
	"github.com/RyanW02/gvm/class/reader"
)

type ClassesEntry struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

func NewClassesEntry() ClassesEntry {
	return ClassesEntry{}
}

func (e *ClassesEntry) Read(cr *reader.ClassReader) (err error) {
	e.InnerClassInfoIndex, err = cr.ReadUint16()
	e.OuterClassInfoIndex, err = cr.ReadUint16()
	e.InnerNameIndex, err = cr.ReadUint16()
	e.InnerClassAccessFlags, err = cr.ReadUint16()
	return
}
