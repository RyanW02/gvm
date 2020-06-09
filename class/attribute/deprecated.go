package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type DeprecatedInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func NewDeprecatedInfo() *DeprecatedInfo {
	return &DeprecatedInfo{}
}

func (i *DeprecatedInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	return
}
