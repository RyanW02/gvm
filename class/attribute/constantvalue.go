package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type ConstantValueInfo struct {
	ConstantValueIndex uint16
}

func NewConstantValueInfo() *ConstantValueInfo {
	return &ConstantValueInfo{}
}

func (i *ConstantValueInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.ConstantValueIndex, err = cr.ReadUint16()
	return
}
