package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type SyntheticInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func NewSyntheticInfo() *SyntheticInfo {
	return &SyntheticInfo{}
}

func (i *SyntheticInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	return
}
