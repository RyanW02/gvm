package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type AnnotationDefaultInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	DefaultValue       ElementValue
}

func NewAnnotationDefaultInfo() *AnnotationDefaultInfo {
	return &AnnotationDefaultInfo{}
}

func (i *AnnotationDefaultInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.DefaultValue = NewElementValue()
	err = i.DefaultValue.Read(cr)
	return
}
