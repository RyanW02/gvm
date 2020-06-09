package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type LocalVariableTableInfo struct {
	AttributeNameIndex       uint16
	AttributeLength          uint32
	LocalVariableTableLength uint16
	LocalVariableTable       []LocalVariableTableEntry
}

func NewLocalVariableTableInfo() *LocalVariableTableInfo {
	return &LocalVariableTableInfo{}
}

func (i *LocalVariableTableInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.LocalVariableTableLength, err = cr.ReadUint16()

	i.LocalVariableTable = make([]LocalVariableTableEntry, i.LocalVariableTableLength)
	for index := uint16(0); index < i.LocalVariableTableLength; index++ {
		i.LocalVariableTable[index] = NewLocalVariableTableEntry()
		err = i.LocalVariableTable[index].Read(cr)
	}

	return
}
