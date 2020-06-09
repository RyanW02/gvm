package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type SignatureInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	SignatureIndex     uint16
}

func NewSignatureInfo() *SignatureInfo {
	return &SignatureInfo{}
}

func (i *SignatureInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.SignatureIndex, err = cr.ReadUint16()
	return
}
