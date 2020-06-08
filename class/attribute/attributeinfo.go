package attribute

import (
	"errors"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
	"github.com/RyanW02/gvm/class/stack"
)

type AttributeInfo interface {
	Read(*reader.ClassReader, constants.ConstantPool) error
}

type ConstantValueInfo struct {
	ConstantValueIndex uint16
}

func NewConstantValueInfo() *ConstantValueInfo {
	return &ConstantValueInfo{}
}

func (i *ConstantValueInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.ConstantValueIndex, err = cr.ReadUint16()
	return
}

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
	for index := uint16(0); index < i.ExceptionTableLength; index++ {
		i.Attributes[index] = NewAttribute()
		err = i.Attributes[index].Read(cr, pool)
	}

	return
}

type StackMapTableInfo struct {
	NumberOfEntries uint16
	Entries         []stack.Frame
}

func NewStackMapTableInfo() *StackMapTableInfo {
	return &StackMapTableInfo{}
}

func (i *StackMapTableInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.NumberOfEntries, err = cr.ReadUint16()
	return
}

var ErrInvalidAttributeName = errors.New("attribute for specified name not registered")

func NewAttributeInfo(name AttributeName) (info AttributeInfo, err error) {
	switch name {
	case ConstantValue:
		info = NewConstantValueInfo()
	case Code:
		info = NewCodeInfo()
	default:
		err = ErrInvalidAttributeName
	}

	return
}
