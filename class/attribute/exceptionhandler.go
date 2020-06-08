package attribute

import "github.com/RyanW02/gvm/class/reader"

type ExceptionHandler struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func NewExceptionHandler() ExceptionHandler {
	return ExceptionHandler{}
}

func (eh *ExceptionHandler) Read(cr *reader.ClassReader) (err error) {
	eh.StartPc, err = cr.ReadUint16()
	eh.EndPc, err = cr.ReadUint16()
	eh.HandlerPc, err = cr.ReadUint16()
	eh.CatchType, err = cr.ReadUint16()
	return
}
