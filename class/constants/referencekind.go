package constants

import "errors"

type ReferenceKind uint8

const (
	RefGetField ReferenceKind = iota + 1
	RefGetStatic
	RefPutField
	RefPutStatic
	RefInvokeVirtual
	RefInvokeStatic
	RefInvokeSpecial
	RefNewInvokeSpecial
	RefInvokeInterface
)

var ErrInvalidReferenceKind = errors.New("reference kind was < 1 or > 9")
