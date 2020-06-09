package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type ExceptionsInfo struct {
	AttributeNameIndex  uint16
	AttributeLength     uint32
	NumberOfExceptions  uint16
	ExceptionIndexTable []uint16
}

func NewExceptionsInfo() *ExceptionsInfo {
	return &ExceptionsInfo{}
}

func (i *ExceptionsInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.NumberOfExceptions, err = cr.ReadUint16()

	i.ExceptionIndexTable = make([]uint16, i.NumberOfExceptions)
	for index := uint16(0); index < i.NumberOfExceptions; index++ {
		i.ExceptionIndexTable[index], err = cr.ReadUint16()
	}

	return
}
