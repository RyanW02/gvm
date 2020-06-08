package constants

import "errors"

type ConstantType uint8

const (
	ConstantClass              ConstantType = 7
	ConstantFieldRef           ConstantType = 9
	ConstantMethodRef          ConstantType = 10
	ConstantInterfaceMethodRef ConstantType = 11
	ConstantString             ConstantType = 8
	ConstantInteger            ConstantType = 3
	ConstantFloat              ConstantType = 4
	ConstantLong               ConstantType = 5
	ConstantDouble             ConstantType = 6
	ConstantNameAndType        ConstantType = 12
	ConstantUtf8               ConstantType = 1
	ConstantMethodHandle       ConstantType = 15
	ConstantMethodType         ConstantType = 16
	ConstantInvokeDynamic      ConstantType = 18
)

var ErrInvalidConstantType = errors.New("no constants found with the specified tag")
