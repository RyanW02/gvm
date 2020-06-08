package stack

import (
	"github.com/RyanW02/gvm/class/reader"
	"github.com/RyanW02/gvm/class/stack/verificationtype"
)

type Frame interface {
	Read(*reader.ClassReader) error
}

type SameFrame struct {
	FrameType FrameType
}

func NewSameFrame() *SameFrame {
	return &SameFrame{
		FrameType: SameFrameType,
	}
}

func (f *SameFrame) Read(cr *reader.ClassReader) (err error) {
	return
}

type SameLocals1StackItemFrame struct {
	FrameType FrameType
	Stack     []verificationtype.VerificationTypeInfo
}

func NewSameLocals1StackItemFrame() *SameLocals1StackItemFrame {
	return &SameLocals1StackItemFrame{
		FrameType: SameLocals1StackItemFrameType,
	}
}

func (f *SameLocals1StackItemFrame) Read(cr *reader.ClassReader) (err error) {

	return
}



