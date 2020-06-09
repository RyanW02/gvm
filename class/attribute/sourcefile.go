package attribute

import (
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type SourceFileInfo struct {
	SourceFileIndex uint16
}

func NewSourceFileInfo() *SourceFileInfo {
	return &SourceFileInfo{}
}

func (i *SourceFileInfo) Read(cr *reader.ClassReader, _ constants.ConstantPool) (err error) {
	i.SourceFileIndex, err = cr.ReadUint16()
	return
}
