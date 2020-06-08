package verificationtype

import (
	"errors"
	"github.com/RyanW02/gvm/class/reader"
)

type VerificationTypeInfo interface {
	Read(*reader.ClassReader) error
}

type TopVariableInfo struct {
	Tag VerificationType
}

func NewTopVariableInfo() *TopVariableInfo {
	return &TopVariableInfo{
		Tag: Top,
	}
}

func (i *TopVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type IntegerVariableInfo struct {
	Tag VerificationType
}

func NewIntegerVariableInfo() *IntegerVariableInfo {
	return &IntegerVariableInfo{
		Tag: Integer,
	}
}

func (i *IntegerVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type FloatVariableInfo struct {
	Tag VerificationType
}

func NewFloatVariableInfo() *FloatVariableInfo {
	return &FloatVariableInfo{
		Tag: Float,
	}
}

func (i *FloatVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type DoubleVariableInfo struct {
	Tag VerificationType
}

func NewDoubleVariableInfo() *DoubleVariableInfo {
	return &DoubleVariableInfo{
		Tag: Double,
	}
}

func (i *DoubleVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type LongVariableInfo struct {
	Tag VerificationType
}

func NewLongVariableInfo() *LongVariableInfo {
	return &LongVariableInfo{
		Tag: Long,
	}
}

func (i *LongVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type NullVariableInfo struct {
	Tag VerificationType
}

func NewNullVariableInfo() *NullVariableInfo {
	return &NullVariableInfo{
		Tag: Null,
	}
}

func (i *NullVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type UninitializedThisVariableInfo struct {
	Tag VerificationType
}

func NewUninitializedThisVariableInfo() *UninitializedThisVariableInfo {
	return &UninitializedThisVariableInfo{
		Tag: UninitializedThis,
	}
}

func (i *UninitializedThisVariableInfo) Read(cr *reader.ClassReader) (err error) {
	return
}

type ObjectVariableInfo struct {
	Tag        VerificationType
	CpoolIndex uint16
}

func NewObjectVariableInfo() *ObjectVariableInfo {
	return &ObjectVariableInfo{
		Tag: Object,
	}
}

func (i *ObjectVariableInfo) Read(cr *reader.ClassReader) (err error) {
	i.CpoolIndex, err = cr.ReadUint16()
	return
}

type UninitializedVariableInfo struct {
	Tag    VerificationType
	Offset uint16
}

func NewUninitializedVariableInfo() *UninitializedVariableInfo {
	return &UninitializedVariableInfo{
		Tag: Uninitialized,
	}
}

func (i *UninitializedVariableInfo) Read(cr *reader.ClassReader) (err error) {
	i.Offset, err = cr.ReadUint16()
	return
}

var ErrInvalidVerificationType = errors.New("no verification type with tag found")

func NewVariableInfo(verificationType VerificationType) (info VerificationTypeInfo, err error) {
	switch verificationType {
	case 0:
		info = NewTopVariableInfo()
	case 1:
		info = NewIntegerVariableInfo()
	case 2:
		info = NewFloatVariableInfo()
	case 3:
		info = NewDoubleVariableInfo()
	case 4:
		info = NewLongVariableInfo()
	case 5:
		info = NewNullVariableInfo()
	case 6:
		info = NewUninitializedThisVariableInfo()
	case 7:
		info = NewObjectVariableInfo()
	case 8:
		info = NewUninitializedVariableInfo()
	default:
		err = ErrInvalidVerificationType
	}

	return
}
