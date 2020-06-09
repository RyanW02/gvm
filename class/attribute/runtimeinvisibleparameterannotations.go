package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type RuntimeInvisibleParameterAnnotationsInfo struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	NumParameters        uint8
	ParameterAnnotations []ParameterAnnotation
}

func NewRuntimeInvisibleParameterAnnotationsInfo() *RuntimeInvisibleParameterAnnotationsInfo {
	return &RuntimeInvisibleParameterAnnotationsInfo{}
}

func (i *RuntimeInvisibleParameterAnnotationsInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.NumParameters, err = cr.ReadUint8()

	i.ParameterAnnotations = make([]ParameterAnnotation, i.NumParameters)
	for index := uint8(0); index < i.NumParameters; index++ {
		i.ParameterAnnotations[index] = NewParameterAnnotation()
		err = i.ParameterAnnotations[index].Read(cr)
	}

	return
}
