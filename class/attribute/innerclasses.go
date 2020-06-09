package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type InnerClassesInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfClasses    uint16
	Classes            []ClassesEntry
}

func NewInnerClassesInfo() *InnerClassesInfo {
	return &InnerClassesInfo{}
}

func (i *InnerClassesInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.NumberOfClasses, err = cr.ReadUint16()

	i.Classes = make([]ClassesEntry, i.NumberOfClasses)
	for index := uint16(0); index < i.NumberOfClasses; index++ {
		i.Classes[index] = NewClassesEntry()
		err = i.Classes[index].Read(cr)
	}

	return
}
