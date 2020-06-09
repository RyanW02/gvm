package attribute

import (
	"github.com/RyanW02/gvm/class/reader"
)

type BootstrapMethod struct {
	BootstrapMethodRef    uint16
	NumBootstrapArguments uint16
	BootstrapArguments    []uint16 // index in constant pool
}

func NewBootstrapMethod() BootstrapMethod {
	return BootstrapMethod{}
}

func (i *BootstrapMethod) Read(cr *reader.ClassReader) (err error) {
	i.BootstrapMethodRef, err = cr.ReadUint16()
	i.NumBootstrapArguments, err = cr.ReadUint16()

	i.BootstrapArguments = make([]uint16, i.NumBootstrapArguments)
	for index := uint16(0); index < i.NumBootstrapArguments; index++ {
		i.BootstrapArguments[index], err = cr.ReadUint16()
	}

	return
}
