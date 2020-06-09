package constants

import (
	"github.com/RyanW02/gvm/class/reader"
)

type ConstantPool []ConstantInfo

func (p ConstantPool) Read(count uint16, cr *reader.ClassReader) (err error) {
	for i := uint16(0); i < count - 1; i++ {
		var tag ConstantType
		{
			typeId, e := cr.ReadUint8()
			if e != nil {
				err = e
			}

			tag = ConstantType(typeId)
		}

		var info ConstantInfo
		info, err = NewConstantInfo(tag)

		err = info.Read(cr)
		p[i] = info
	}

	return
}

func (p ConstantPool) GetUtf8(index uint16) *ConstantUtf8Info {
	return p[index - 1].(*ConstantUtf8Info)
}
