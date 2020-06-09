package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type LocalVariableTypeTableInfo struct {
	AttributeNameIndex           uint16
	AttributeLength              uint32
	LocalVariableTypeTableLength uint16
	LocalVariableTable           []LocalVariableTypeTableEntry
}

func NewLocalVariableTypeTableInfo() *LocalVariableTypeTableInfo {
	return &LocalVariableTypeTableInfo{}
}

func (i *LocalVariableTypeTableInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.LocalVariableTypeTableLength, err = cr.ReadUint16()

	i.LocalVariableTable = make([]LocalVariableTypeTableEntry, i.LocalVariableTypeTableLength)
	for index := uint16(0); index < i.LocalVariableTypeTableLength; index++ {
		i.LocalVariableTable[index] = NewLocalVariableTypeTableEntry()
		err = i.LocalVariableTable[index].Read(cr)
	}

	return
}
