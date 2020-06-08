package verificationtype

type VerificationType uint8

const (
	Top VerificationType = iota
	Integer
	Float
	Double
	Long
	Null
	UninitializedThis
	Object
	Uninitialized
)
