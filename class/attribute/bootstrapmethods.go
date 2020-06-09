package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type BootstrapMethodsInfo struct {
	AttributeNameIndex  uint16
	AttributeLength     uint32
	NumBootstrapMethods uint16
	BootstrapMethods    []BootstrapMethod
}

func NewBootstrapMethodsInfo() *BootstrapMethodsInfo {
	return &BootstrapMethodsInfo{}
}

func (i *BootstrapMethodsInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.NumBootstrapMethods, err = cr.ReadUint16()

	i.BootstrapMethods = make([]BootstrapMethod, i.NumBootstrapMethods)
	for index := uint16(0); index < i.NumBootstrapMethods; index++ {
		i.BootstrapMethods[index] = NewBootstrapMethod()
		err = i.BootstrapMethods[index].Read(cr)
	}

	return
}
