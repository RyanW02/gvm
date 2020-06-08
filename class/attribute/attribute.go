package attribute

import (
	"fmt"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type Attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               AttributeInfo
}

func NewAttribute() Attribute {
	return Attribute{}
}

func (a *Attribute) Read(cr *reader.ClassReader, pool constants.ConstantPool) (err error) {
	a.AttributeNameIndex, err = cr.ReadUint16()
	a.AttributeLength, err = cr.ReadUint32()
	a.Info, err = NewAttributeInfo(a.GetAttributeName(pool))
	fmt.Println(a.GetAttributeName(pool))
	err = a.Info.Read(cr, pool)
	return
}

func (a *Attribute) GetAttributeName(pool constants.ConstantPool) AttributeName {
	return AttributeName(pool.GetUtf8(a.AttributeNameIndex).String())
}
