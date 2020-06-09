package attribute

import "github.com/RyanW02/gvm/class/reader"

type ElementValue struct {
	Tag   byte
	Value ElementValueValue
}

func NewElementValue() ElementValue {
	return ElementValue{}
}

func (v *ElementValue) Read(cr *reader.ClassReader) (err error) {
	v.Tag, err = cr.ReadUint8()
	v.Value, err = NewElementValueValue(v.Tag)
	return
}
