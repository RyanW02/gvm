package stack

import (
	"github.com/RyanW02/gvm/class/reader"
	"github.com/RyanW02/gvm/class/stack/verificationtype"
)

// https://github.com/openjdk/jdk11u/blob/master/src/hotspot/share/classfile/stackMapTable.cpp#L212

type Frame interface {
	Read(cr *reader.ClassReader, previous Frame) error
	GetOffset() uint16
}

func GetFrame(tag uint8) (frame Frame, err error) {
	frameType, err := GetFrameType(tag)
	if err != nil {
		return
	}

	switch frameType {
	case SameFrameType:
		frame = NewSameFrame(tag)
	case SameLocals1StackItemFrameType:
		frame = NewSameLocals1StackItemFrame(tag)
	case SameLocals1StackItemFrameExtendedType:
		frame = NewSameLocals1StackItemFrameExtended()
	case ChopFrameType:
		frame = NewChopFrame(tag)
	case SameFrameExtendedType:
		frame = NewSameFrameExtended()
	case AppendFrameType:
		frame = NewAppendFrame(tag)
	case FullFrameType:
		frame = NewFullFrame()
	default:
		err = ErrInvalidFrameType
	}

	return
}

type SameFrame struct {
	FrameType FrameType
	tag       uint16
	Offset    uint16
}

func NewSameFrame(tag uint8) *SameFrame {
	return &SameFrame{
		FrameType: SameFrameType,
		tag:       uint16(tag),
	}
}

func (f *SameFrame) GetOffset() uint16 {
	return f.Offset
}

func (f *SameFrame) Read(cr *reader.ClassReader, previous Frame) (err error) {
	if previous == nil {
		f.Offset = f.tag
	} else {
		f.Offset = previous.GetOffset() + f.tag + 1
	}

	return
}

type SameLocals1StackItemFrame struct {
	FrameType FrameType
	tag       uint16
	Offset    uint16
	Stack     []verificationtype.VerificationTypeInfo
}

func NewSameLocals1StackItemFrame(tag uint8) *SameLocals1StackItemFrame {
	return &SameLocals1StackItemFrame{
		FrameType: SameLocals1StackItemFrameType,
		tag:       uint16(tag),
		Stack:     make([]verificationtype.VerificationTypeInfo, 1),
	}
}

func (f *SameLocals1StackItemFrame) GetOffset() uint16 {
	return f.Offset
}

func (f *SameLocals1StackItemFrame) Read(cr *reader.ClassReader, previous Frame) (err error) {
	if previous == nil {
		f.Offset = f.tag - 64
	} else {
		f.Offset = previous.GetOffset() + f.tag - 63
	}

	var tag uint8
	tag, err = cr.ReadUint8()
	verificationType := verificationtype.VerificationType(tag)

	f.Stack[0], err = verificationtype.NewVerificationTypeInfo(verificationType)
	err = f.Stack[0].Read(cr)

	return
}

type SameLocals1StackItemFrameExtended struct {
	FrameType FrameType
	Offset    uint16
	Stack     []verificationtype.VerificationTypeInfo
}

func NewSameLocals1StackItemFrameExtended() *SameLocals1StackItemFrameExtended {
	return &SameLocals1StackItemFrameExtended{
		FrameType: SameLocals1StackItemFrameExtendedType,
		Stack:     make([]verificationtype.VerificationTypeInfo, 1),
	}
}

func (f *SameLocals1StackItemFrameExtended) GetOffset() uint16 {
	return f.Offset
}

func (f *SameLocals1StackItemFrameExtended) Read(cr *reader.ClassReader, previous Frame) (err error) {
	var delta uint16
	delta, err = cr.ReadUint16()

	if previous == nil {
		f.Offset = delta
	} else {
		f.Offset = previous.GetOffset() + delta + 1
	}

	var tag uint8
	tag, err = cr.ReadUint8()
	verificationType := verificationtype.VerificationType(tag)

	f.Stack[0], err = verificationtype.NewVerificationTypeInfo(verificationType)
	err = f.Stack[0].Read(cr)
	return
}

type ChopFrame struct {
	FrameType FrameType
	tag       uint8
	Offset    uint16
}

func NewChopFrame(tag uint8) *ChopFrame {
	return &ChopFrame{
		FrameType: ChopFrameType,
		tag:       tag,
	}
}

func (f *ChopFrame) GetOffset() uint16 {
	return f.Offset
}

func (f *ChopFrame) Read(cr *reader.ClassReader, previous Frame) (err error) {
	var delta uint16
	delta, err = cr.ReadUint16()

	if previous == nil {
		f.Offset = delta
	} else {
		f.Offset = previous.GetOffset() + delta + 1
	}

	return
}

// k last locals are missing
func (f *ChopFrame) MissingLocals() uint8 {
	return f.tag - 251
}

type SameFrameExtended struct {
	FrameType FrameType
	Offset    uint16
}

func NewSameFrameExtended() *ChopFrame {
	return &ChopFrame{
		FrameType: SameFrameExtendedType,
	}
}

func (f *SameFrameExtended) GetOffset() uint16 {
	return f.Offset
}

func (f *SameFrameExtended) Read(cr *reader.ClassReader, previous Frame) (err error) {
	var delta uint16
	delta, err = cr.ReadUint16()

	if previous == nil {
		f.Offset = delta
	} else {
		f.Offset = previous.GetOffset() + delta + 1
	}

	return
}

type AppendFrame struct {
	FrameType FrameType
	tag       uint8
	Offset    uint16
	Locals    []verificationtype.VerificationTypeInfo
}

func NewAppendFrame(tag uint8) *AppendFrame {
	return &AppendFrame{
		FrameType: AppendFrameType,
		tag:       tag,
	}
}

func (f *AppendFrame) GetOffset() uint16 {
	return f.Offset
}

func (f *AppendFrame) Read(cr *reader.ClassReader, previous Frame) (err error) {
	var delta uint16
	delta, err = cr.ReadUint16()

	if previous == nil {
		f.Offset = delta
	} else {
		f.Offset = previous.GetOffset() + delta + 1
	}

	f.Locals = make([]verificationtype.VerificationTypeInfo, f.tag-251)
	for i := 0; i < len(f.Locals); i++ {
		var tag uint8
		tag, err = cr.ReadUint8()
		verificationType := verificationtype.VerificationType(tag)

		var info verificationtype.VerificationTypeInfo
		info, err = verificationtype.NewVerificationTypeInfo(verificationType)
		err = info.Read(cr)

		f.Locals[i] = info
	}

	return
}

type FullFrame struct {
	FrameType          FrameType
	Offset             uint16
	NumberOfLocals     uint16
	Locals             []verificationtype.VerificationTypeInfo
	NumberOfStackItems uint16
	Stack              []verificationtype.VerificationTypeInfo
}

func NewFullFrame() *FullFrame {
	return &FullFrame{
		FrameType: FullFrameType,
	}
}

func (f *FullFrame) GetOffset() uint16 {
	return f.Offset
}

func (f *FullFrame) Read(cr *reader.ClassReader, previous Frame) (err error) {
	var delta uint16
	delta, err = cr.ReadUint16()

	if previous == nil {
		f.Offset = delta
	} else {
		f.Offset = previous.GetOffset() + delta + 1
	}

	f.NumberOfLocals, err = cr.ReadUint16()

	// read locals
	f.Locals = make([]verificationtype.VerificationTypeInfo, f.NumberOfLocals)
	for i := uint16(0); i < f.NumberOfLocals; i++ {
		var tag uint8
		tag, err = cr.ReadUint8()
		verificationType := verificationtype.VerificationType(tag)

		var info verificationtype.VerificationTypeInfo
		info, err = verificationtype.NewVerificationTypeInfo(verificationType)
		err = info.Read(cr)

		f.Locals[i] = info
	}

	f.NumberOfStackItems, err = cr.ReadUint16()

	// read stack
	f.Stack = make([]verificationtype.VerificationTypeInfo, f.NumberOfStackItems)
	for i := uint16(0); i < f.NumberOfStackItems; i++ {
		var tag uint8
		tag, err = cr.ReadUint8()
		verificationType := verificationtype.VerificationType(tag)

		var info verificationtype.VerificationTypeInfo
		info, err = verificationtype.NewVerificationTypeInfo(verificationType)
		err = info.Read(cr)

		f.Stack[i] = info
	}

	return
}
