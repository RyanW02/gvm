package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type SourceDebugExtensionInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	DebugExtension     []byte
}

func NewSourceDebugExtensionInfo() *SourceDebugExtensionInfo {
	return &SourceDebugExtensionInfo{}
}

func (i *SourceDebugExtensionInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.AttributeNameIndex, err = cr.ReadUint16()
	i.AttributeLength, err = cr.ReadUint32()
	i.DebugExtension, err = cr.Read(int(i.AttributeLength))
	return
}
