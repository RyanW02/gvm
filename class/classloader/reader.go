package classloader

import (
	"github.com/RyanW02/gvm/class"
	"github.com/RyanW02/gvm/class/attribute"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/field"
	"github.com/RyanW02/gvm/class/field/method"
	"github.com/RyanW02/gvm/class/reader"
)

func (cl *ClassLoader) Read(cr *reader.ClassReader) (class class.Class, err error) {
	class.Magic, err = cr.ReadUint32()
	if class.Magic != reader.Magic {
		err = reader.ErrInvalidMagic
	}

	class.MinorVersion, err = cr.ReadUint16()
	class.MajorVersion, err = cr.ReadUint16()
	class.ConstantPoolCount, err = cr.ReadUint16()

	class.ConstantPool = make(constants.ConstantPool, class.ConstantPoolCount - 1)
	err = class.ConstantPool.Read(class.ConstantPoolCount, cr)

	class.AccessFlags, err = cr.ReadUint16()

	class.ThisClass, err = cr.ReadUint16()
	class.SuperClass, err = cr.ReadUint16()
	class.InterfacesCount, err = cr.ReadUint16()

	// Read interfaces
	class.Interfaces = make([]uint16, class.InterfacesCount)
	for i := uint16(0); i < class.InterfacesCount; i++ {
		class.Interfaces[i], err = cr.ReadUint16()
	}

	class.FieldsCount, err = cr.ReadUint16()

	// Read fields
	class.Fields = make([]field.Field, class.FieldsCount)
	for i := uint16(0); i < class.FieldsCount; i++ {
		field := field.NewField()
		err = field.Read(cr, class.ConstantPool)
		class.Fields[i] = field
	}

	class.MethodsCount, err = cr.ReadUint16()

	// Read methods
	class.Methods = make([]method.Method, class.MethodsCount)
	for i := uint16(0); i < class.MethodsCount; i++ {
		method := method.NewMethod()
		err = method.Read(cr, class.ConstantPool)
		class.Methods[i] = method
	}

	class.AttributesCount, err = cr.ReadUint16()

	// Read methods
	class.Attributes = make([]attribute.Attribute, class.AttributesCount)
	for i := uint16(0); i < class.AttributesCount; i++ {
		attribute := attribute.NewAttribute()
		err = attribute.Read(cr, class.ConstantPool)
		class.Attributes[i] = attribute
	}

	return class, err
}

