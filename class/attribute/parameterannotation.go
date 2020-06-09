package attribute

import "github.com/RyanW02/gvm/class/reader"

type ParameterAnnotation struct {
	NumAnnotations uint16
	Annotations    []Annotation
}

func NewParameterAnnotation() ParameterAnnotation {
	return ParameterAnnotation{}
}

func (a *ParameterAnnotation) Read(cr *reader.ClassReader) (err error) {
	a.NumAnnotations, err = cr.ReadUint16()

	a.Annotations = make([]Annotation, a.NumAnnotations)
	for i := uint16(0); i < a.NumAnnotations; i++ {
		a.Annotations[i] = NewAnnotation()
		err = a.Annotations[i].Read(cr)
	}

	return
}
