package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type EnclosingMethodInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ClassIndex         uint16
	MethodIndex        uint16
}

func NewEnclosingMethodInfo() *EnclosingMethodInfo {
	return &EnclosingMethodInfo{}
}

func (i *EnclosingMethodInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.ClassIndex, err = cr.ReadUint16()
	i.MethodIndex, err = cr.ReadUint16()
	return
}
