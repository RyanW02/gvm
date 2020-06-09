package attribute

import (
	"errors"
	"github.com/RyanW02/gvm/class/constants"
	"github.com/RyanW02/gvm/class/reader"
)

type AttributeInfo interface {
	Read(*reader.ClassReader, constants.ConstantPool) error
}

var ErrInvalidAttributeName = errors.New("attribute for specified name not registered")

func NewAttributeInfo(name AttributeName) (info AttributeInfo, err error) {
	switch name {
	case ConstantValue:
		info = NewConstantValueInfo()
	case Code:
		info = NewCodeInfo()
	case StackMapTable:
		info = NewStackMapTableInfo()
	case Exceptions:
		info = NewExceptionsInfo()
	case InnerClasses:
		info = NewInnerClassesInfo()
	case EnclosingMethod:
		info = NewEnclosingMethodInfo()
	case Synthetic:
		info = NewSyntheticInfo()
	case Signature:
		info = NewSignatureInfo()
	case SourceFile:
		info = NewSourceFileInfo()
	case SourceDebugExtension:
		info = NewSourceDebugExtensionInfo()
	case LineNumberTable:
		info = NewLineNumberTableInfo()
	case LocalVariableTable:
		info = NewLocalVariableTableInfo()
	case LocalVariableTypeTable:
		info = NewLocalVariableTypeTableInfo()
	case Deprecated:
		info = NewDeprecatedInfo()
	case RuntimeVisibleAnnotations:
		info = NewRuntimeVisibleAnnotationsInfo()
	case RuntimeInvisibleAnnotations:
		info = NewRuntimeInvisibleAnnotationsInfo()
	case RuntimeInvisibleParameterAnnotations:
		info = NewRuntimeInvisibleParameterAnnotationsInfo()
	case AnnotationDefault:
		info = NewAnnotationDefaultInfo()
	case BootstrapMethods:
		info = NewBootstrapMethodsInfo()
	default:
		err = ErrInvalidAttributeName
	}

	return
}
