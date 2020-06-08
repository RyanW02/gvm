package constants

import (
	"github.com/RyanW02/gvm/class/reader"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4
type ConstantInfo interface {
	Read(*reader.ClassReader) error
}

type ConstantClassInfo struct {
	Tag       ConstantType
	NameIndex uint16
}

func NewConstantClassInfo() *ConstantClassInfo {
	return &ConstantClassInfo{
		Tag: ConstantClass,
	}
}

func (c *ConstantClassInfo) Read(cr *reader.ClassReader) (err error) {
	c.NameIndex, err = cr.ReadUint16()
	return
}

type ConstantFieldRefInfo struct {
	Tag              ConstantType
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func NewConstantFieldRefInfo() *ConstantFieldRefInfo {
	return &ConstantFieldRefInfo{
		Tag: ConstantFieldRef,
	}
}

func (c *ConstantFieldRefInfo) Read(cr *reader.ClassReader) (err error) {
	c.ClassIndex, err = cr.ReadUint16()
	c.NameAndTypeIndex, err = cr.ReadUint16()
	return
}

type ConstantMethodRefInfo struct {
	Tag              ConstantType
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func NewConstantMethodRefInfo() *ConstantMethodRefInfo {
	return &ConstantMethodRefInfo{
		Tag: ConstantMethodRef,
	}
}

func (c *ConstantMethodRefInfo) Read(cr *reader.ClassReader) (err error) {
	c.ClassIndex, err = cr.ReadUint16()
	c.NameAndTypeIndex, err = cr.ReadUint16()
	return
}

type ConstantInterfaceMethodRefInfo struct {
	Tag              ConstantType
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func NewConstantInterfaceMethodRefInfo() *ConstantInterfaceMethodRefInfo {
	return &ConstantInterfaceMethodRefInfo{
		Tag: ConstantInterfaceMethodRef,
	}
}

func (c *ConstantInterfaceMethodRefInfo) Read(cr *reader.ClassReader) (err error) {
	c.ClassIndex, err = cr.ReadUint16()
	c.NameAndTypeIndex, err = cr.ReadUint16()
	return
}

type ConstantStringInfo struct {
	Tag         ConstantType
	StringIndex uint16
}

func NewConstantStringInfo() *ConstantStringInfo {
	return &ConstantStringInfo{
		Tag: ConstantString,
	}
}

func (c *ConstantStringInfo) Read(cr *reader.ClassReader) (err error) {
	c.StringIndex, err = cr.ReadUint16()
	return
}

type ConstantIntegerInfo struct {
	Tag   ConstantType
	Bytes uint32
}

func NewConstantIntegerInfo() *ConstantIntegerInfo {
	return &ConstantIntegerInfo{
		Tag: ConstantInteger,
	}
}

func (c *ConstantIntegerInfo) Read(cr *reader.ClassReader) (err error) {
	c.Bytes, err = cr.ReadUint32()
	return
}

type ConstantFloatInfo struct {
	Tag   ConstantType
	Bytes uint32
}

func NewConstantFloatInfo() *ConstantFloatInfo {
	return &ConstantFloatInfo{
		Tag: ConstantFloat,
	}
}

func (c *ConstantFloatInfo) Read(cr *reader.ClassReader) (err error) {
	c.Bytes, err = cr.ReadUint32()
	return
}

type ConstantLongInfo struct {
	Tag       ConstantType
	HighBytes uint32
	LowBytes  uint32
}

func NewConstantLongInfo() *ConstantLongInfo {
	return &ConstantLongInfo{
		Tag: ConstantLong,
	}
}

func (c *ConstantLongInfo) Read(cr *reader.ClassReader) (err error) {
	c.HighBytes, err = cr.ReadUint32()
	c.LowBytes, err = cr.ReadUint32()
	return
}

type ConstantDoubleInfo struct {
	Tag       ConstantType
	HighBytes uint32
	LowBytes  uint32
}

func NewConstantDoubleInfo() *ConstantDoubleInfo {
	return &ConstantDoubleInfo{
		Tag: ConstantDouble,
	}
}

func (c *ConstantDoubleInfo) Read(cr *reader.ClassReader) (err error) {
	c.HighBytes, err = cr.ReadUint32()
	c.LowBytes, err = cr.ReadUint32()
	return
}

type ConstantNameAndTypeInfo struct {
	Tag             ConstantType
	NameIndex       uint16
	DescriptorIndex uint16
}

func NewConstantNameAndTypeInfo() *ConstantNameAndTypeInfo {
	return &ConstantNameAndTypeInfo{
		Tag: ConstantNameAndType,
	}
}

func (c *ConstantNameAndTypeInfo) Read(cr *reader.ClassReader) (err error) {
	c.NameIndex, err = cr.ReadUint16()
	c.DescriptorIndex, err = cr.ReadUint16()
	return
}

type ConstantUtf8Info struct {
	Tag    ConstantType
	Length uint16
	Bytes  []byte
}

func NewConstantUtf8Info() *ConstantUtf8Info {
	return &ConstantUtf8Info{
		Tag: ConstantUtf8,
	}
}

func (c *ConstantUtf8Info) Read(cr *reader.ClassReader) (err error) {
	c.Length, err = cr.ReadUint16()
	c.Bytes, err = cr.Read(int(c.Length))
	return
}

// helper func
func (c *ConstantUtf8Info) String() string {
	return string(c.Bytes)
}

type ConstantMethodHandleInfo struct {
	Tag            ConstantType
	ReferenceKind  ReferenceKind
	ReferenceIndex uint16
}

func NewConstantMethodHandleInfo() *ConstantMethodHandleInfo {
	return &ConstantMethodHandleInfo{
		Tag: ConstantMethodHandle,
	}
}

func (c *ConstantMethodHandleInfo) Read(cr *reader.ClassReader) (err error) {
	kind, err := cr.ReadUint8()
	if err == nil && (kind > 9 || kind < 1) {
		err = ErrInvalidReferenceKind
	}

	c.ReferenceKind = ReferenceKind(kind)
	c.ReferenceIndex, err = cr.ReadUint16()
	return
}

type ConstantMethodTypeInfo struct {
	Tag             ConstantType
	DescriptorIndex uint16
}

func NewConstantMethodTypeInfo() *ConstantMethodTypeInfo {
	return &ConstantMethodTypeInfo{
		Tag: ConstantMethodType,
	}
}

func (c *ConstantMethodTypeInfo) Read(cr *reader.ClassReader) (err error) {
	c.DescriptorIndex, err = cr.ReadUint16()
	return
}

type ConstantInvokeDynamicInfo struct {
	Tag                      ConstantType
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func NewConstantInvokeDynamicInfo() *ConstantInvokeDynamicInfo {
	return &ConstantInvokeDynamicInfo{
		Tag: ConstantInvokeDynamic,
	}
}

func (c *ConstantInvokeDynamicInfo) Read(cr *reader.ClassReader) (err error) {
	c.BootstrapMethodAttrIndex, err = cr.ReadUint16()
	c.NameAndTypeIndex, err = cr.ReadUint16()
	return
}

func NewConstantInfo(tag ConstantType) (info ConstantInfo, err error) {
	switch tag {
	case ConstantClass:
		info = NewConstantClassInfo()
	case ConstantFieldRef:
		info = NewConstantFieldRefInfo()
	case ConstantMethodRef:
		info = NewConstantMethodRefInfo()
	case ConstantInterfaceMethodRef:
		info = NewConstantMethodRefInfo()
	case ConstantString:
		info = NewConstantStringInfo()
	case ConstantInteger:
		info = NewConstantIntegerInfo()
	case ConstantFloat:
		info = NewConstantFloatInfo()
	case ConstantLong:
		info = NewConstantLongInfo()
	case ConstantDouble:
		info = NewConstantDoubleInfo()
	case ConstantNameAndType:
		info = NewConstantNameAndTypeInfo()
	case ConstantUtf8:
		info = NewConstantUtf8Info()
	case ConstantMethodHandle:
		info = NewConstantMethodHandleInfo()
	case ConstantMethodType:
		info = NewConstantMethodTypeInfo()
	case ConstantInvokeDynamic:
		info = NewConstantInvokeDynamicInfo()
	default:
		err = ErrInvalidConstantType
	}

	return
}
