package attribute

type AttributeName string

const (
	ConstantValue                        AttributeName = "ConstantValue"
	Code                                 AttributeName = "Code"
	StackMapTable                        AttributeName = "StackMapTable"
	Exceptions                           AttributeName = "Exceptions"
	InnerClasses                         AttributeName = "InnerClasses"
	EnclosingMethod                      AttributeName = "EnclosingMethod"
	Synthetic                            AttributeName = "Synthetic"
	Signature                            AttributeName = "Signature"
	SourceFile                           AttributeName = "SourceFile"
	SourceDebugExtension                 AttributeName = "SourceDebugExtension"
	LineNumberTable                      AttributeName = "LineNumberTable"
	LocalVariableTable                   AttributeName = "LocalVariableTable"
	LocalVariableTypeTable               AttributeName = "LocalVariableTypeTable"
	Deprecated                           AttributeName = "Deprecated"
	RuntimeVisibleAnnotations            AttributeName = "RuntimeVisibleAnnotations"
	RuntimeInvisibleAnnotations          AttributeName = "RuntimeInvisibleAnnotations"
	RuntimeVisibleParameterAnnotations   AttributeName = "RuntimeVisibleParameterAnnotations"
	RuntimeInvisibleParameterAnnotations AttributeName = "RuntimeInvisibleParameterAnnotations"
	AnnotationDefault                    AttributeName = "AnnotationDefault"
	BootstrapMethods                     AttributeName = "BootstrapMethods"
)
