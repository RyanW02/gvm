package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type RuntimeVisibleAnnotationsInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumAnnotations     uint16
	Annotations        []Annotation
}

func NewRuntimeVisibleAnnotationsInfo() *RuntimeVisibleAnnotationsInfo {
	return &RuntimeVisibleAnnotationsInfo{}
}

func (i *RuntimeVisibleAnnotationsInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.NumAnnotations, err = cr.ReadUint16()

	i.Annotations = make([]Annotation, i.NumAnnotations)
	for index := uint16(0); index < i.NumAnnotations; index++ {
		i.Annotations[index] = NewAnnotation()
		err = i.Annotations[index].Read(cr)
	}

	return
}
