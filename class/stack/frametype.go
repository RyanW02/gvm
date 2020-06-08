package stack

import "errors"

// integer value doesn't represent anything
type FrameType struct {
	Min uint8 // inclusive
	Max uint8 // inclusive
}

var (
	SameFrameType = FrameType{
		Min: 0,
		Max: 63,
	}
	SameLocals1StackItemFrameType = FrameType{
		Min: 64,
		Max: 127,
	}
	SameLocals1StackItemFrameExtendedType = FrameType{
		Min: 247,
		Max: 247,
	}
	ChopFrameType = FrameType{
		Min: 248,
		Max: 250,
	}
	SameFrameExtendedType = FrameType{
		Min: 251,
		Max: 251,
	}
	AppendFrameType = FrameType{
		Min: 252,
		Max: 254,
	}
	FullFrameType = FrameType{
		Min: 255,
		Max: 255,
	}

	allTypes = []FrameType{
		SameFrameType,
		SameLocals1StackItemFrameType,
		SameLocals1StackItemFrameExtendedType,
		ChopFrameType,
		SameFrameExtendedType,
		AppendFrameType,
		FullFrameType,
	}
)

var ErrInvalidFrameType = errors.New("a frame with the specified tag was not found")

func GetFrameType(tag uint8) (frame FrameType, err error) {
	var found bool
	for _, frameType := range allTypes {
		if tag >= frameType.Min && tag <= frameType.Max {
			frame = frameType
			found = true
			break
		}
	}

	if !found {
		err = ErrInvalidFrameType
	}

	return
}
