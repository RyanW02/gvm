package field

import (
	"github.com/RyanW02/gvm/class/access"
	"github.com/RyanW02/gvm/class/attribute"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type Field struct {
	AccessFlags     []access.AccessFlag
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []attribute.Attribute
}

func NewField() Field {
	return Field{}
}

func (f *Field) Read(cr *reader.ClassReader, pool constants.ConstantPool) (err error) {
	var flags uint16
	flags, err = cr.ReadUint16()
	f.AccessFlags = access.GetFlags(flags)

	f.NameIndex, err = cr.ReadUint16()
	f.DescriptorIndex, err = cr.ReadUint16()
	f.AttributesCount, err = cr.ReadUint16()

	f.Attributes = make([]attribute.Attribute, f.AttributesCount)
	for i := uint16(0); i < f.AttributesCount; i++ {
		attribute := attribute.NewAttribute()
		err = attribute.Read(cr, pool)
		f.Attributes[i] = attribute
	}

	return
}
