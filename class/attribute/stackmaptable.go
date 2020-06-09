package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
	"github.com/RyanW02/gvm/class/stack"
)

type StackMapTableInfo struct {
	NumberOfEntries uint16
	Entries         []stack.Frame
}

func NewStackMapTableInfo() *StackMapTableInfo {
	return &StackMapTableInfo{}
}

func (i *StackMapTableInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.NumberOfEntries, err = cr.ReadUint16()

	i.Entries = make([]stack.Frame, i.NumberOfEntries)
	var previous stack.Frame
	for index := uint16(0); index < i.NumberOfEntries; index++ {
		var tag uint8
		tag, err = cr.ReadUint8()
		var frame stack.Frame

		frame, err = stack.GetFrame(tag)
		err = frame.Read(cr, previous)
		i.Entries[index] = frame

		previous = frame
	}

	return
}
