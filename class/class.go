package class

import (
	"github.com/RyanW02/gvm/class/attribute"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/field"
	"github.com/RyanW02/gvm/class/field/method"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
type Class struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      constants.ConstantPool // The constant_pool table is indexed from 1 to constant_pool_count-1
	AccessFlags       uint16
	ThisClass         uint16 // Index in constants pool
	SuperClass        uint16 // 0 or index in constants pool
	InterfacesCount   uint16
	Interfaces        []uint16 // Index in constants pool
	FieldsCount       uint16
	Fields            []field.Field
	MethodsCount      uint16
	Methods           []method.Method
	AttributesCount   uint16
	Attributes        []attribute.Attribute
}
