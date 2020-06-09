package attribute

import "github.com/RyanW02/gvm/class/reader"

type ElementValuePair struct {
	ElementNameIndex uint16
	Value            ElementValue
}

func NewElementValuePair() ElementValuePair {
	return ElementValuePair{}
}

func (p *ElementValuePair) Read(cr *reader.ClassReader) (err error) {
	p.ElementNameIndex, err = cr.ReadUint16()
	p.Value = NewElementValue()
	err = p.Value.Read(cr)
	return
}
