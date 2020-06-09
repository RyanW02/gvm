package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type LineNumberTableInfo struct {
	LineNumberTableLength uint16
	LineNumberTable       []LineNumberTableEntry
}

func NewLineNumberTableInfo() *LineNumberTableInfo {
	return &LineNumberTableInfo{}
}

func (i *LineNumberTableInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.LineNumberTableLength, err = cr.ReadUint16()

	i.LineNumberTable = make([]LineNumberTableEntry, i.LineNumberTableLength)
	for index := uint16(0); index < i.LineNumberTableLength; index++ {
		i.LineNumberTable[index] = NewLineNumberTableEntry()
		err = i.LineNumberTable[index].Read(cr)
	}

	return
}
