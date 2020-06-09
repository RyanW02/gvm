package attribute

import "github.com/RyanW02/gvm/class/reader"

type Annotation struct {
	TypeIndex            uint16
	NumElementValuePairs uint16
	ElementValuePairs    []ElementValuePair
}

func NewAnnotation() Annotation {
	return Annotation{}
}

func (a *Annotation) Read(cr *reader.ClassReader) (err error) {
	a.TypeIndex, err = cr.ReadUint16()
	a.NumElementValuePairs, err = cr.ReadUint16()

	a.ElementValuePairs = make([]ElementValuePair, a.NumElementValuePairs)
	for i := uint16(0); i < a.NumElementValuePairs; i++ {
		a.ElementValuePairs[i] = NewElementValuePair()
		err = a.ElementValuePairs[i].Read(cr)
	}

	return
}
