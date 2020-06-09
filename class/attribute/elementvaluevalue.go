package attribute

import (
	"errors"
	"github.com/RyanW02/gvm/class/reader"
)

type ElementValueValue interface {
	Read(*reader.ClassReader) error
}

var ErrInvalidElementValueValue = errors.New("invalid element_value value tag")

func NewElementValueValue(tag byte) (value ElementValueValue, err error) {
	switch tag {
	case byte('B'):
		value = NewElementConstValue()
	case byte('C'):
		value = NewElementConstValue()
	case byte('D'):
		value = NewElementConstValue()
	case byte('F'):
		value = NewElementConstValue()
	case byte('I'):
		value = NewElementConstValue()
	case byte('J'):
		value = NewElementConstValue()
	case byte('S'):
		value = NewElementConstValue()
	case byte('Z'):
		value = NewElementConstValue()
	case byte('s'):
		value = NewElementConstValue()
	case byte('e'):
		value = NewElementEnumConstValue()
	case byte('@'):
		value = NewElementAnnotationValue()
	case byte('['):
		value = NewElementArrayValue()
	default:
		err = ErrInvalidElementValueValue
	}

	return
}

type ElementConstValue struct {
	ConstValueIndex uint16
}

func NewElementConstValue() *ElementConstValue {
	return &ElementConstValue{}
}

func (v *ElementConstValue) Read(cr *reader.ClassReader) (err error) {
	v.ConstValueIndex, err = cr.ReadUint16()
	return
}

type ElementEnumConstValue struct {
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

func NewElementEnumConstValue() *ElementEnumConstValue {
	return &ElementEnumConstValue{}
}

func (v *ElementEnumConstValue) Read(cr *reader.ClassReader) (err error) {
	v.TypeNameIndex, err = cr.ReadUint16()
	v.ConstNameIndex, err = cr.ReadUint16()
	return
}

type ElementAnnotationValue struct {
	AnnotationValue Annotation
}

func NewElementAnnotationValue() *ElementAnnotationValue {
	return &ElementAnnotationValue{}
}

func (v *ElementAnnotationValue) Read(cr *reader.ClassReader) error {
	v.AnnotationValue = NewAnnotation()
	return v.AnnotationValue.Read(cr)
}

type ElementArrayValue struct {
	NumValues uint16
	Values    []ElementValue
}

func NewElementArrayValue() *ElementArrayValue {
	return &ElementArrayValue{}
}

func (v *ElementArrayValue) Read(cr *reader.ClassReader) (err error) {
	v.NumValues, err = cr.ReadUint16()

	v.Values = make([]ElementValue, v.NumValues)
	for i := uint16(0); i < v.NumValues; i++ {
		v.Values[i] = NewElementValue()
		err = v.Values[i].Read(cr)
	}

	return
}
