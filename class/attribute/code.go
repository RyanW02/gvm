package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type CodeInfo struct {
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionHandler
	AttributesCount      uint16
	Attributes           []Attribute
}

func NewCodeInfo() *CodeInfo {
	return &CodeInfo{}
}

func (i *CodeInfo) Read(cr *reader.ClassReader, pool constants.ConstantPool) (err error) {
	i.MaxStack, err = cr.ReadUint16()
	i.MaxLocals, err = cr.ReadUint16()
	i.CodeLength, err = cr.ReadUint32()
	i.Code, err = cr.Read(int(i.CodeLength))
	i.ExceptionTableLength, err = cr.ReadUint16()

	i.ExceptionTable = make([]ExceptionHandler, i.ExceptionTableLength)
	for index := uint16(0); index < i.ExceptionTableLength; index++ {
		i.ExceptionTable[index] = NewExceptionHandler()
		err = i.ExceptionTable[index].Read(cr)
	}

	i.AttributesCount, err = cr.ReadUint16()

	i.Attributes = make([]Attribute, i.AttributesCount)
	for index := uint16(0); index < i.AttributesCount; index++ {
		i.Attributes[index] = NewAttribute()
		err = i.Attributes[index].Read(cr, pool)
	}

	return
}
